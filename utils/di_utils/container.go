package di_utils

import (
	"context"
	"os"
	"runtime/pprof"
	"time"

	"github.com/teadove/teasutils/utils/context_utils"

	"golang.org/x/sync/errgroup"

	"github.com/teadove/teasutils/utils/notify_utils"
	"github.com/teadove/teasutils/utils/perf_utils"
	"github.com/teadove/teasutils/utils/settings_utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/refrect_utils"
)

type CloserWithContext interface {
	Close(ctx context.Context) error
}

type Health interface {
	Health(ctx context.Context) error
}

type Container interface {
	Healths() []Health
	Closers() []CloserWithContext
}

func withProfiler(ctx context.Context) error {
	file, err := os.Create(settings_utils.BaseSettings.Prof.ResultFilename)
	if err != nil {
		return errors.Wrap(err, "could not open result file")
	}

	err = pprof.StartCPUProfile(file)
	if err != nil {
		return errors.Wrap(err, "could not start CPU profile")
	}

	zerolog.Ctx(ctx).Warn().Msg("cpu.profile.started")

	return nil
}

func stop(ctx context.Context, container Container) error {
	errorsGroup, ctx := errgroup.WithContext(ctx)

	ctx, cancel := context.WithTimeout(ctx, settings_utils.BaseSettings.Metrics.CloseTimeout)
	defer cancel()

	for _, stoper := range container.Closers() {
		errorsGroup.Go(func() error {
			return context_utils.CPUCancel(ctx, stoper.Close)
		})
	}

	err := errorsGroup.Wait()
	if err != nil {
		return errors.Wrap(err, "could not stop container")
	}

	return nil
}

func BuildFromSettings[T Container](
	ctx context.Context,
	builder func(ctx context.Context) (T, error),
) (T, error) {
	if settings_utils.BaseSettings.Prof.SpamMemUsage {
		go perf_utils.SpamLogMemUsage(ctx, settings_utils.BaseSettings.Prof.SpamMemUsagePeriod)
		zerolog.Ctx(ctx).
			Warn().
			Str("period", settings_utils.BaseSettings.Prof.SpamMemUsagePeriod.String()).
			Msg("spam.memory.usage.added")
	}

	if settings_utils.BaseSettings.Prof.Enabled {
		err := withProfiler(ctx)
		if err != nil {
			return *new(T), errors.Wrap(err, "failed to add profiler")
		}
	}

	t0 := time.Now()

	builtContainer, err := builder(ctx)
	if err != nil {
		return *new(T), errors.Wrap(err, "build container failed")
	}

	if !settings_utils.BaseSettings.Release {
		err = checkFromCheckers(ctx, builtContainer.Healths())
		if err != nil {
			return *new(T), errors.Wrap(err, "health check failed")
		}
	}

	runMetricsFromSettingsInBackground(ctx, builtContainer)
	notify_utils.RunOnInterruptAndExit(func() {
		t0 = time.Now()

		zerolog.Ctx(ctx).
			Debug().
			Msg("stopping.container")

		if settings_utils.BaseSettings.Prof.Enabled {
			pprof.StopCPUProfile()
		}

		err = stop(ctx, builtContainer)
		if err != nil {
			zerolog.Ctx(ctx).
				Error().
				Str("elapsed", time.Since(t0).String()).
				Err(err).Stack().
				Msg("could not stop container")
		}

		zerolog.Ctx(ctx).
			Info().
			Str("elapsed", time.Since(t0).String()).
			Msg("container.stoped")
	})

	zerolog.Ctx(ctx).
		Info().
		Str("container", refrect_utils.GetTypesStringRepresentation(builtContainer)).
		Str("elapsed", time.Since(t0).String()).
		Msg("container.built")

	return builtContainer, nil
}

func MustBuildFromSettings[T Container](
	ctx context.Context,
	builder func(ctx context.Context) (T, error),
) T {
	t, err := BuildFromSettings[T](ctx, builder)
	if err != nil {
		panic(errors.Wrap(err, "build container failed"))
	}

	return t
}

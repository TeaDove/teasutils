package di_utils

import (
	"context"
	"os"
	"runtime/pprof"
	"time"

	"github.com/teadove/teasutils/utils/reflect_utils"

	"github.com/teadove/teasutils/utils/context_utils"

	"golang.org/x/sync/errgroup"

	"github.com/teadove/teasutils/utils/notify_utils"
	"github.com/teadove/teasutils/utils/perf_utils"
	"github.com/teadove/teasutils/utils/settings_utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func withProfiler(ctx context.Context) error {
	file, err := os.Create(settings_utils.ServiceSettings.Prof.ResultFilename)
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

func stop(ctx context.Context, stoppers []any) error {
	errorsGroup, ctx := errgroup.WithContext(ctx)

	ctx, cancel := context.WithTimeout(ctx, settings_utils.ServiceSettings.Metrics.CloseTimeout)
	defer cancel()

	for _, stopper := range stoppers {
		v := reflect_utils.ConvertToWithCtxAndErr(stopper)
		if v == nil {
			zerolog.Ctx(ctx).
				Error().
				Str("stopper", reflect_utils.GetFunctionName(v)).
				Msg("stopper.is.null")
		}

		errorsGroup.Go(func() error {
			return context_utils.CPUCancel(ctx, v)
		})
	}

	err := errorsGroup.Wait()
	if err != nil {
		return errors.Wrap(err, "could not stop container")
	}

	return nil
}

func BuildFromSettings[T any](
	ctx context.Context,
	builder func(ctx context.Context) (T, []any, []any, error),
) (T, error) {
	if settings_utils.ServiceSettings.Prof.SpamMemUsage {
		go perf_utils.SpamLogMemUsage(ctx, settings_utils.ServiceSettings.Prof.SpamMemUsagePeriod)
		zerolog.Ctx(ctx).
			Warn().
			Str("period", settings_utils.ServiceSettings.Prof.SpamMemUsagePeriod.String()).
			Msg("spam.memory.usage.added")
	}

	if settings_utils.ServiceSettings.Prof.Enabled {
		err := withProfiler(ctx)
		if err != nil {
			return *new(T), errors.Wrap(err, "failed to add profiler")
		}
	}

	t0 := time.Now()

	builtContainer, healthers, stoppers, err := builder(ctx)
	if err != nil {
		return *new(T), errors.Wrap(err, "build container failed")
	}

	if !settings_utils.ServiceSettings.Release {
		err = checkFromCheckers(ctx, healthers)
		if err != nil {
			return *new(T), errors.Wrap(err, "health check failed")
		}
	}

	runMetricsFromSettingsInBackground(ctx, healthers)
	notify_utils.RunOnInterruptAndExit(func() {
		t0 = time.Now()

		zerolog.Ctx(ctx).
			Debug().
			Msg("stopping.container")

		if settings_utils.ServiceSettings.Prof.Enabled {
			pprof.StopCPUProfile()
		}

		err = stop(ctx, stoppers)
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
			Msg("container.stopped")
	})

	zerolog.Ctx(ctx).
		Info().
		Str("container", reflect_utils.GetTypesStringRepresentation(builtContainer)).
		Str("elapsed", time.Since(t0).String()).
		Msg("container.built")

	return builtContainer, nil
}

func MustBuildFromSettings[T any](
	ctx context.Context,
	builder func(ctx context.Context) (T, []any, []any, error),
) T {
	t, err := BuildFromSettings[T](ctx, builder)
	if err != nil {
		panic(errors.Wrap(err, "build container failed"))
	}

	return t
}

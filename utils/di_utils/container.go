package di_utils

import (
	"context"
	"os"
	"os/signal"
	"runtime/pprof"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/refrect_utils"
)

type Container interface {
	Health(ctx context.Context) []error
}

func withProfiler(ctx context.Context) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			pprof.StopCPUProfile()
			os.Exit(0)
		}
	}()

	zerolog.Ctx(ctx).Error().Msg("cpu.profile.started")

	return nil
}

func BuildFromSettings[T Container](
	ctx context.Context,
	builder func(ctx context.Context) (T, error),
) (T, error) {
	t0 := time.Now()

	builtContainer, err := builder(ctx)
	if err != nil {
		return *new(T), errors.Wrap(err, "build container failed")
	}

	runMetricsFromSettingsInBackground(ctx, builtContainer)

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

package di_utils

import (
	"context"

	"github.com/teadove/teasutils/utils/context_utils"
	"github.com/teadove/teasutils/utils/settings_utils"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/rs/zerolog"
)

func checkFromCheckers(ctx context.Context, checkers []Health) error {
	ctx, cancel := context.WithTimeout(ctx, settings_utils.BaseSettings.Metrics.RequestTimeout)
	defer cancel()

	var healthCheckable Health

	errGroup, ctx := errgroup.WithContext(ctx)
	for _, healthCheckable = range checkers {
		errGroup.Go(func() error { return context_utils.CPUCancel(ctx, healthCheckable.Health) })
	}

	err := errGroup.Wait()
	if err != nil {
		err = errors.Wrap(err, "failed to check")
		zerolog.Ctx(ctx).
			Error().Stack().
			Err(err).
			Msg("checking.failed")

		return err
	}

	return nil
}

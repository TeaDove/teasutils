package di_utils

import (
	"context"

	"github.com/teadove/teasutils/utils/context_utils"
	"github.com/teadove/teasutils/utils/reflect_utils"
	"github.com/teadove/teasutils/utils/settings_utils"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/rs/zerolog"
)

func checkFromCheckers(ctx context.Context, checkers []any) error {
	ctx, cancel := context.WithTimeout(ctx, settings_utils.ServiceSettings.Metrics.RequestTimeout)
	defer cancel()

	errGroup, ctx := errgroup.WithContext(ctx)

	for _, checker := range checkers {
		v := reflect_utils.ConvertToWithCtxAndErr(checker)
		if v == nil {
			zerolog.Ctx(ctx).
				Error().
				Str("health", reflect_utils.GetFunctionName(v)).
				Msg("health.check.is.null")
		}

		errGroup.Go(func() error { return context_utils.CPUCancel(ctx, v) })
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

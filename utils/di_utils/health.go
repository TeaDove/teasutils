package di_utils

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/rs/zerolog"
)

func checkFromCheckers(ctx context.Context, checkers []func(ctx context.Context) error) error {
	var checker func(ctx context.Context) error

	errGroup, ctx := errgroup.WithContext(ctx)
	for _, checker = range checkers {
		errGroup.Go(func() error { return checker(ctx) })
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

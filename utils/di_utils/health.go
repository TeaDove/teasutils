package di_utils

import (
	"context"
	stderrors "errors"

	"github.com/rs/zerolog"
)

func checkFromCheckers(ctx context.Context, checkers []func(ctx context.Context) error) []error {
	var (
		errors  = make([]error, 0)
		err     error
		checker func(ctx context.Context) error
	)

	for _, checker = range checkers {
		err = checker(ctx)
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		zerolog.Ctx(ctx).
			Error().Stack().
			Err(stderrors.Join(errors...)).
			Msg("checking.failed")
	}

	return errors
}

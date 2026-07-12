package mustutils

import (
	"context"

	"github.com/cockroachdb/errors"

	"github.com/teadove/teasutils/utils/reflectutils"

	"github.com/rs/zerolog"
)

func WithRecover(f func(ctx context.Context) error) func(ctx context.Context) error {
	return func(ctx context.Context) (err error) {
		defer func() {
			recovered := AnyToErr(recover())
			if recovered != nil {
				err = errors.Wrap(recovered, "recovered")
			}
		}()

		return f(ctx)
	}
}

// WithRecoverAndLog
// Decorates function with panic recovering and log for errors.
func WithRecoverAndLog(f func(ctx context.Context) error) func(ctx context.Context) {
	return func(ctx context.Context) {
		err := WithRecover(f)(ctx)
		if err != nil {
			zerolog.Ctx(ctx).
				Error().
				Stack().
				Err(err).
				Str("func_name", reflectutils.GetFunctionName(f)).
				Msg("failed.to.execute.function")
		}
	}
}

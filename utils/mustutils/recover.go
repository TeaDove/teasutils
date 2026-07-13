package mustutils

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"

	"github.com/teadove/teasutils/utils/reflectutils"

	"github.com/rs/zerolog"
)

// WithRecover decorates f so that a panic is recovered and returned as an
// error wrapped with "recovered"; a normal error from f is returned as-is.
func WithRecover(f func() error) func() error {
	return func() (err error) {
		defer func() {
			recovered := AnyToErr(recover())
			if recovered != nil {
				err = errors.Wrap(recovered, "recovered")
			}
		}()

		return f()
	}
}

// WithRecoverAndLog decorates f with panic recovery and logs any resulting
// error (from a panic or a returned error) via the ctx logger.
func WithRecoverAndLog(f func(ctx context.Context) error) func(ctx context.Context) {
	return func(ctx context.Context) {
		err := WithRecover(func() error { return f(ctx) })()
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

// DoInBackground runs fn in a new goroutine, recovering panics and logging
// errors. The goroutine's context keeps ctx's values but is detached from its
// cancellation and capped with a 3-minute timeout, so fn outlives the caller.
func DoInBackground(ctx context.Context, fn func(ctx context.Context) error) {
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 3*time.Minute)

	go func() {
		defer cancel()

		err := WithRecover(func() error { return fn(ctx) })()
		if err != nil {
			zerolog.Ctx(ctx).
				Error().
				Stack().
				Err(err).
				Str("func_name", reflectutils.GetFunctionName(fn)).
				Msg("failed.to.execute.function")
		}
	}()
}

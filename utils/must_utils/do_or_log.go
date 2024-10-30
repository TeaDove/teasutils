package must_utils

import (
	"context"

	"github.com/teadove/teasutils/utils/refrect_utils"

	"github.com/rs/zerolog"
)

func DoOrLog(
	f func(ctx context.Context) error,
	errorMsg string,
) func(ctx context.Context) {
	return func(ctx context.Context) {
		err := f(ctx)
		if err != nil {
			zerolog.Ctx(ctx).
				Error().
				Err(err).
				Str("func_name", refrect_utils.GetFunctionName(f)).
				Msg(errorMsg)
		}
	}
}

func DoOrLogWithStacktrace(
	f func(ctx context.Context) error,
	errorMsg string,
) func(ctx context.Context) {
	return func(ctx context.Context) {
		err := f(ctx)
		if err != nil {
			zerolog.Ctx(ctx).
				Error().
				Stack().
				Err(err).
				Str("func_name", refrect_utils.GetFunctionName(f)).
				Msg(errorMsg)
		}
	}
}

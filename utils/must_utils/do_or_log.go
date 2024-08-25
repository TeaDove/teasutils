package must_utils

import (
	"context"

	"teasutils/utils/refrect_utils"

	"github.com/rs/zerolog"
)

func DoOrLog(f func(ctx context.Context) error) func(ctx context.Context) {
	return func(ctx context.Context) {
		err := f(ctx)
		if err != nil {
			zerolog.Ctx(ctx).
				Error().
				Err(err).
				Str("func_name", refrect_utils.GetFunctionName(f)).
				Msg("do.failed")
		}
	}
}

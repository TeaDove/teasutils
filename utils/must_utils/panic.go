package must_utils

import (
	"context"

	"github.com/rs/zerolog"
)

func FancyPanic(ctx context.Context, err error) {
	zerolog.Ctx(ctx).Panic().Stack().Err(err).Msg("panicked")
}

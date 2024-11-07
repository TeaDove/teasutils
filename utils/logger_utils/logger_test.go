package logger_utils

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestUnit_LoggerUtils_ErrorWithStackrace_Ok(t *testing.T) {
	t.Parallel()

	err := errors.WithStack(errors.New("test error"))
	ctx := NewLoggedCtx()

	zerolog.Ctx(ctx).Error().Stack().Err(err).Msg("error")
}

func TestUnit_LoggerUtils_ErrorWithStackraceInJson_Ok(t *testing.T) {
	t.Parallel()

	logger := makeLogger(&settings{
		Level:   "DEBUG",
		Factory: "JSON",
	})

	err := errors.WithStack(errors.New("test error"))
	ctx := NewLoggedCtx()

	ctx = WithStrContextLog(ctx, "userId", "123")

	ctx = logger.WithContext(ctx)

	zerolog.Ctx(ctx).Error().Stack().Err(err).Msg("error")
}

func TestUnit_LoggerUtils_Panic_Ok(t *testing.T) {
	t.Parallel()

	err := errors.WithStack(errors.New("test error"))
	ctx := NewLoggedCtx()

	assert.Panics(t, func() {
		zerolog.Ctx(ctx).Panic().Stack().Err(err).Msg("error")
	})
}

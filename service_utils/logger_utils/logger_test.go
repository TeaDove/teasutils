package logger_utils

import (
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestUnit_LoggerUtils_ErrorWithStackrace_Ok(_ *testing.T) { //nolint: paralleltest // racing
	err := errors.WithStack(errors.New("test error"))
	ctx := NewLoggedCtx()

	zerolog.Ctx(ctx).Error().Stack().Err(err).Msg("error")
	zerolog.Ctx(ctx).Error().Err(err).Msg("error")
}

func TestUnit_LoggerUtils_ErrorWithStackraceInJson_Ok(_ *testing.T) { //nolint: paralleltest // racing
	logger := makeLogger("debug", "json")

	err := errors.WithStack(errors.New("test error"))
	ctx := NewLoggedCtx()

	ctx = WithValue(ctx, "userId", "123")

	ctx = logger.WithContext(ctx)

	zerolog.Ctx(ctx).Error().Stack().Err(err).Msg("error")
}

func TestUnit_LoggerUtils_Panic_Ok(t *testing.T) { //nolint: paralleltest // racing
	err := errors.WithStack(errors.New("test error"))
	ctx := NewLoggedCtx()

	assert.Panics(t, func() {
		zerolog.Ctx(ctx).Panic().Stack().Err(err).Msg("error")
	})
}

func TestUnit_LoggerUtils_ReadWriteCtx_Ok(t *testing.T) { //nolint: paralleltest // racing
	ctx := NewLoggedCtx()

	ctx = WithReadableValue(ctx, "userId", "123")
	act := ReadValue(ctx, "userId")
	assert.Equal(t, "123", act)

	ctx = WithReadableValue(ctx, "appId", "123")
	act = ReadValue(ctx, "appIdWrong")
	assert.Empty(t, act)

	act = ReadValue(ctx, "somethingOther")
	assert.Empty(t, act)
}

func TestUnit_LoggerUtils_WithValues(t *testing.T) {
	t.Parallel()

	ctx := NewLoggedCtx()
	ctx = WithValue(ctx, "userId", "123", "username", "peter")

	zerolog.Ctx(ctx).Info().Msg("log")

}

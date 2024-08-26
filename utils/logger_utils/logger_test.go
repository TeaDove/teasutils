package logger_utils

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestUnit_LoggerUtils_ErrorWithStackrace_Ok(t *testing.T) {
	err := errors.WithStack(errors.New("test error"))
	ctx := NewLoggedCtx()

	zerolog.Ctx(ctx).Error().Stack().Err(err).Msg("error")
}

func TestUnit_LoggerUtils_Panic_Ok(t *testing.T) {
	err := errors.WithStack(errors.New("test error"))
	ctx := NewLoggedCtx()

	assert.Panics(t, func() {
		zerolog.Ctx(ctx).Panic().Stack().Err(err).Msg("error")
	})
}

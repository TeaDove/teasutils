package logger_utils

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/teadove/teasutils/utils/must_utils"

	"github.com/teadove/teasutils/utils/settings_utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func NewLoggedCtx() context.Context {
	return AddLoggerToCtx(context.Background())
}

func AddLoggerToCtx(ctx context.Context) context.Context {
	return globalLogger.With().Logger().WithContext(ctx)
}

func WithStrContextLog(ctx context.Context, key string, value string) context.Context {
	return zerolog.Ctx(ctx).With().Str(key, value).Ctx(ctx).Logger().WithContext(ctx)
}

func humanMarshalStack(err error) any {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	e, ok := err.(stackTracer)
	if !ok {
		return nil
	}

	stack := e.StackTrace()
	formatted := ""

	for _, frame := range stack {
		formatted += fmt.Sprintf("%+v\n", frame)
	}

	return formatted
}

func makeLogger() zerolog.Logger {
	//nolint: reassign // Need this
	zerolog.ErrorStackMarshaler = humanMarshalStack

	level := must_utils.Must(zerolog.ParseLevel(settings_utils.BaseSettings.Log.Level))

	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Caller().
		Logger().
		Level(level)

	if strings.EqualFold(settings_utils.BaseSettings.Log.Factory, "CONSOLE") {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	logger.Trace().Str("logLevel", level.String()).Msg("logger.initiated")

	return logger
}

//nolint:gochecknoglobals // need this
var globalLogger = makeLogger()

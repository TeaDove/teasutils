package logger_utils

import (
	"context"
	"fmt"
	"github.com/teadove/teasutils/utils/must_utils"
	"os"

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

func initLogger(settings *settings) {
	zerolog.ErrorStackMarshaler = humanMarshalStack

	level := must_utils.MustNoCtx(zerolog.ParseLevel(settings.LogLevel))

	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Caller().
		Logger().
		Level(level)

	if settings.LoggerFactory == "console" {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	globalLogger = logger
	logger.Trace().Str("log.level", level.String()).Msg("logger.initiated")
}

var globalLogger = zerolog.Logger{}

func init() {
	type BaseSettings struct {
		Logger settings `envPrefix:"log__"`
	}

	baseSettings, err := settings_utils.InitSetting[BaseSettings](context.Background())
	if err != nil {
		panic(errors.Wrap(err, "failed to init base settings"))
	}

	initLogger(&baseSettings.Logger)
}

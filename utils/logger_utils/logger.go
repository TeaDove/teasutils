package logger_utils

import (
	"context"
	"fmt"
	"os"

	"github.com/teadove/teasutils/utils/settings_utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

func humanMarshalStack(err error) interface{} {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}
	e, ok := err.(stackTracer)
	if !ok {
		return nil
	}
	// It's mean when env=dev just print track
	for _, frame := range e.StackTrace() {
		fmt.Printf("%+s:%d\r\n", frame, frame)
	}
	return nil
}

func initLogger(settings *settings) {
	zerolog.ErrorStackMarshaler = humanMarshalStack

	level, err := zerolog.ParseLevel(settings.LogLevel)
	if err != nil {
		log.Error().
			Stack().
			Err(err).
			Str("decision", "debug.will.be.used").
			Msg("invalid.log.level")

		level = zerolog.DebugLevel
	}

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

var globalLogger = zerolog.New(os.Stdout).With().Timestamp().Logger()

func init() {
	type BaseSettings struct {
		Logger settings `envPrefix:"logger__"`
	}

	baseSettings, err := settings_utils.InitSetting[BaseSettings](context.Background())
	if err != nil {
		panic(errors.Wrap(err, "failed to init base settings"))
	}

	initLogger(&baseSettings.Logger)
}

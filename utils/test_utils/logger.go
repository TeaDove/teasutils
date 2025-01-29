package test_utils

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

func GetLoggedContext() context.Context {
	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Caller().
		Logger().
		Level(zerolog.DebugLevel)
	logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	return logger.With().Logger().WithContext(context.Background())
}

package loggerutils

import (
	"fmt"
	"os"
	"strings"

	"github.com/teadove/teasutils/serviceutils/settingsutils"

	"github.com/teadove/teasutils/utils/mustutils"

	"github.com/rs/zerolog"
)

func printedMarshalStack(err error) any {
	fmt.Printf("%+v\n", err) //nolint: forbidigo // Allowed for logs

	return "up"
}

func marshalStack(err error) any {
	return fmt.Sprintf("%+v", err)
}

func makeLoggerFromSettings() zerolog.Logger {
	return makeLogger(settingsutils.ServiceSettings.Log.Level, settingsutils.ServiceSettings.Log.Factory)
}

func makeLogger(level, factory string) zerolog.Logger {
	loggerLevel := mustutils.Must(zerolog.ParseLevel(level))

	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Caller().
		Logger().
		Level(loggerLevel)

	if strings.EqualFold(factory, "CONSOLE") {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
		//nolint: reassign // TODO find better solution
		zerolog.ErrorStackMarshaler = printedMarshalStack
	} else {
		//nolint: reassign // TODO find better solution
		zerolog.ErrorStackMarshaler = marshalStack
	}

	logger.Trace().Str("logLevel", loggerLevel.String()).Msg("logger.initiated")

	return logger
}

//nolint:gochecknoglobals // need this
var globalLogger = makeLoggerFromSettings()

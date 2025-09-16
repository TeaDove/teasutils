package logger_utils

import (
	"fmt"
	"github.com/teadove/teasutils/service_utils/settings_utils"
	"os"
	"strings"

	"github.com/teadove/teasutils/utils/must_utils"

	"github.com/rs/zerolog"
)

func printedMarshalStack(err error) any {
	fmt.Printf("%+v\n", err) //nolint: forbidigo // Allowed for logs

	return "up"
}

func marshalStack(err error) any {
	return fmt.Sprintf("%+v", err)
}

func makeLogger() zerolog.Logger {
	level := must_utils.Must(zerolog.ParseLevel(settings_utils.ServiceSettings.Log.Level))

	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Caller().
		Logger().
		Level(level)

	logger.Hook()

	if strings.EqualFold(settings_utils.ServiceSettings.Log.Factory, "CONSOLE") {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
		//nolint: reassign // TODO find better solution
		zerolog.ErrorStackMarshaler = printedMarshalStack
	} else {
		//nolint: reassign // TODO find better solution
		zerolog.ErrorStackMarshaler = marshalStack
	}

	logger.Trace().Str("logLevel", level.String()).Msg("logger.initiated")

	return logger
}

//nolint:gochecknoglobals // need this
var globalLogger = makeLogger()

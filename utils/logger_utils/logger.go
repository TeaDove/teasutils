package logger_utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/teadove/teasutils/utils/must_utils"

	"github.com/teadove/teasutils/utils/settings_utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func printedMarshalStack(err error) any {
	err = errors.WithStack(err)

	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	e, ok := err.(stackTracer)
	if !ok {
		return nil
	}

	stack := e.StackTrace()

	for _, frame := range stack {
		//nolint: forbidigo // only exception
		fmt.Printf("%+v\n", frame)
	}

	return "up"
}

func jsonMarshalStack(err error) any {
	err = errors.WithStack(err)

	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	e, ok := err.(stackTracer)
	if !ok {
		return nil
	}

	stack := e.StackTrace()
	v := strings.Builder{}

	for _, frame := range stack {
		v.WriteString(fmt.Sprintf("%+v\n", frame))
	}

	return v.String()
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
		zerolog.ErrorStackMarshaler = jsonMarshalStack
	}

	logger.Trace().Str("logLevel", level.String()).Msg("logger.initiated")

	return logger
}

//nolint:gochecknoglobals // need this
var globalLogger = makeLogger()

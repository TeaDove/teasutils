package test_utils

import (
	"context"
	"os"

	"github.com/teadove/teasutils/utils/reflect_utils"

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

// LogAny
//
// Logs everything, only for debug purposes.
func LogAny(values ...any) {
	arr := zerolog.Arr()
	for _, value := range values {
		arr.Dict(
			zerolog.Dict().
				Interface(reflect_utils.GetTypesStringRepresentation(value), value),
		)
	}

	zerolog.Ctx(GetLoggedContext()).Info().Array("items", arr).Msg("logging.any")
}

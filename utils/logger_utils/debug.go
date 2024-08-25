package logger_utils

import (
	"github.com/teadove/teasutils/utils/refrect_utils"

	"github.com/rs/zerolog"
)

// LogAny
//
// Logs everything, only for debug purposes
func LogAny(values ...any) {
	arr := zerolog.Arr()
	for _, value := range values {
		arr.Dict(
			zerolog.Dict().
				Interface(refrect_utils.GetTypesStringRepresentation(value), value),
		)
	}

	globalLogger.Info().Array("items", arr).Msg("logging.any")
}

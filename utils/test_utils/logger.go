package test_utils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

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

// Pprint
//
// Prints everything with types, only for debug purposes.
func Pprint(values ...any) {
	var v strings.Builder
	for _, value := range values {
		v.WriteString(color.RedString(fmt.Sprintf("%T", value)))
		v.WriteString(": ")

		encoded, err := json.MarshalIndent(value, "", "  ")
		if err != nil {
			v.WriteString(fmt.Sprintf("%+v", value))
			continue
		}

		v.Write(encoded)
		v.WriteByte('\n')
	}

	println(v.String()) //nolint: forbidigo // required
}

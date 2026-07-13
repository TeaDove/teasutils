package testutils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/rs/zerolog"
)

// Context returns a background context carrying a debug-level zerolog console
// logger. Intended for tests that exercise code reading the logger from ctx.
func Context() context.Context {
	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Caller().
		Logger().
		Level(zerolog.DebugLevel)
	logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	return logger.With().Logger().WithContext(context.Background())
}

// Debug pretty-prints each value with its type as indented JSON to stdout.
// For quick debugging only; not for production logging.
func Debug(values ...any) {
	var v strings.Builder
	v.WriteString("------\n")

	for idx, value := range values {
		fmt.Fprintf(&v, "%d (%s): ", idx, color.RedString(fmt.Sprintf("%T", value)))

		encoded, err := json.MarshalIndent(value, "", "  ")
		if err != nil {
			fmt.Fprintf(&v, "%+v", value)

			continue
		}

		v.Write(encoded)
		v.WriteByte('\n')
	}

	v.WriteString("\n------")
	println(v.String()) //nolint: forbidigo // required
}

package redact_utils

import (
	"context"
	"fmt"
	"slices"

	"github.com/rs/zerolog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const (
	redactedText   = "REDACTED"
	redactedMaxLen = 5
	trimmedMaxLen  = 10
)

func Trim(s string) string {
	if len(s) <= trimmedMaxLen {
		return fmt.Sprintf("[%s]", s)
	}

	return fmt.Sprintf("[%s...:%d]", s[:redactedMaxLen], len(s))
}

func Redact(s string) string {
	if len(s) == 0 {
		return fmt.Sprintf("[%s]", redactedText)
	}

	return fmt.Sprintf("[%s:%d]", redactedText, len(s))
}

func RedactWithPrefix(s string) string {
	if len(s) <= redactedMaxLen {
		return Redact(s)
	}

	return fmt.Sprintf("[%s:%s...:%d]", redactedText, s[:redactedMaxLen], len(s))
}

func RedactJSONWithPrefix(ctx context.Context, s []byte, paths ...string) []byte {
	var (
		err      error
		redacted = slices.Clone(s)
	)

	for _, path := range paths {
		redacted, err = sjson.SetBytes(redacted, path, RedactWithPrefix(gjson.GetBytes(redacted, path).String()))
		if err != nil {
			zerolog.Ctx(ctx).
				Warn().
				Str("s", string(s)).
				Str("path", path).
				Stack().Err(err).
				Msg("failed.to.set.json.path")
		}
	}

	return redacted
}

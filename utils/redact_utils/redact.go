package redact_utils

import "fmt"

const (
	redacted       = "[REDACTED]"
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
		return redacted
	}

	return fmt.Sprintf("[REDACTED:%d]", len(s))
}

func RedactWithPrefix(s string) string {
	if len(s) <= redactedMaxLen {
		return Redact(s)
	}

	return fmt.Sprintf("[REDACTED:%s...:%d]", s[:redactedMaxLen], len(s))
}

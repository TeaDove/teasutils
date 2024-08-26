package redact_utils

import "fmt"

const (
	redacted = "[REDACTED]"
	maxLen   = 5
)

func Redact(s string) string {
	if len(s) == 0 {
		return redacted
	}

	return fmt.Sprintf("[REDACTED:%d]", len(s))
}

func RedactWithPrefix(s string) string {
	if len(s) <= maxLen {
		return Redact(s)
	}

	return fmt.Sprintf("[REDACTED:%s...:%d]", s[:maxLen], len(s))
}

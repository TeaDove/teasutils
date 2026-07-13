package redactutils

import (
	"fmt"
)

const (
	redactedText   = "REDACTED"
	redactedMaxLen = 5
	trimmedMaxLen  = 10
)

// Trim shortens s for logging using the default limit (10 chars), keeping a
// prefix and the original length. Values are not hidden; use Redact for secrets.
func Trim(s string) string {
	return TrimSized(s, trimmedMaxLen)
}

// TrimSized is Trim with an explicit prefix length: it returns "[s]" when s
// fits in maxLen, otherwise "[prefix...:len]".
func TrimSized(s string, maxLen int) string {
	if len(s) <= maxLen {
		return fmt.Sprintf("[%s]", s)
	}

	return fmt.Sprintf("[%s...:%d]", s[:maxLen], len(s))
}

// Redact fully hides s, revealing only that it is redacted and its length,
// e.g. "[REDACTED:12]" (or "[REDACTED]" when empty).
func Redact(s string) string {
	if len(s) == 0 {
		return fmt.Sprintf("[%s]", redactedText)
	}

	return fmt.Sprintf("[%s:%d]", redactedText, len(s))
}

// RedactWithPrefix redacts s but keeps a short readable prefix (5 chars),
// useful for correlating values without exposing them.
func RedactWithPrefix(s string) string {
	return RedactWithPrefixSized(s, redactedMaxLen)
}

// RedactWithPrefixSized is RedactWithPrefix with an explicit prefix length;
// values no longer than maxLen are fully redacted via Redact.
func RedactWithPrefixSized(s string, maxLen int) string {
	if len(s) <= maxLen {
		return Redact(s)
	}

	return fmt.Sprintf("[%s:%s...:%d]", redactedText, s[:maxLen], len(s))
}

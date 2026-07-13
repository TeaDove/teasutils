package randomutils

import (
	"math/rand/v2"
	"strings"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Text returns a non-cryptographic random 16-character alphanumeric string.
func Text() string {
	const stringLen = 16
	return TextWithLen(stringLen)
}

// TextWithLen returns a non-cryptographic random alphanumeric string of the
// given length, or "" if length <= 0. Not safe for secrets or tokens.
func TextWithLen(length int) string {
	if length <= 0 {
		return ""
	}

	var builder strings.Builder
	for range length {
		//nolint: gosec // no need to be secure
		builder.WriteByte(alphabet[rand.IntN(len(alphabet))])
	}

	return builder.String()
}

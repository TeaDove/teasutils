package random_utils

import (
	"math/rand/v2"
	"strings"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Text() string {
	const stringLen = 16
	return TextWithLen(stringLen)
}

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

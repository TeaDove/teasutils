package random_utils

import (
	crypto_rand "crypto/rand"
	"math/rand/v2"
	"strings"
)

const base32alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

func String() string {
	const stringLen = 10
	return StringWithLen(stringLen)
}

func StringWithLen(length int) string {
	if length <= 0 {
		return ""
	}

	var builder strings.Builder
	for range length {
		//nolint: gosec // no need to be secure
		builder.WriteByte(base32alphabet[rand.IntN(len(base32alphabet))])
	}

	return builder.String()
}

func CryptoString() string {
	const stringLen = 10
	return CryptoStringWithLen(stringLen)
}

func CryptoStringWithLen(length int) string {
	src := make([]byte, length)

	_, _ = crypto_rand.Read(src)

	for i := range src {
		src[i] = base32alphabet[src[i]%32]
	}

	return string(src)
}

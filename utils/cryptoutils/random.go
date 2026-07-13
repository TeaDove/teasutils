package cryptoutils

import "crypto/rand"

const base32alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

// Text returns a cryptographically-random 16-character base32 string.
func Text() string {
	const stringLen = 16
	return TextWithLen(stringLen)
}

// TextWithLen returns a cryptographically-random base32 string of the given
// length (the mapping is unbiased, as 256 is a multiple of the alphabet size).
func TextWithLen(length int) string {
	src := make([]byte, length)

	_, _ = rand.Read(src)

	for i := range src {
		src[i] = base32alphabet[src[i]%32]
	}

	return string(src)
}

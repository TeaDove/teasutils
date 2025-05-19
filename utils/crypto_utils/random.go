package crypto_utils

import "crypto/rand"

const base32alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

func Text() string {
	const stringLen = 16
	return TextWithLen(stringLen)
}

func TextWithLen(length int) string {
	src := make([]byte, length)

	_, _ = rand.Read(src)

	for i := range src {
		src[i] = base32alphabet[src[i]%32]
	}

	return string(src)
}

package crypto_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_RandomUtils_CryptoString_Ok(t *testing.T) {
	t.Parallel()

	assert.Len(t, Text(), 16)
}

func TestUnit_RandomUtils_CryptoStringWithLen_Ok(t *testing.T) {
	t.Parallel()

	assert.Len(t, TextWithLen(100), 100)
}

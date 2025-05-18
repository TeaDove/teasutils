package random_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_RandomUtils_String_Ok(t *testing.T) {
	t.Parallel()

	assert.Len(t, String(), 10)
}

func TestUnit_RandomUtils_StringWithLen_Ok(t *testing.T) {
	t.Parallel()

	assert.Len(t, StringWithLen(100), 100)
}

func TestUnit_RandomUtils_CryptoString_Ok(t *testing.T) {
	t.Parallel()

	assert.Len(t, CryptoString(), 10)
}

func TestUnit_RandomUtils_CryptoStringWithLen_Ok(t *testing.T) {
	t.Parallel()

	assert.Len(t, CryptoStringWithLen(100), 100)
}

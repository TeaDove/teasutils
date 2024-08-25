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

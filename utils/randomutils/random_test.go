package randomutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_RandomUtils_String_Ok(t *testing.T) {
	t.Parallel()

	assert.Len(t, Text(), 16)
}

func TestUnit_RandomUtils_StringWithLen_Ok(t *testing.T) {
	t.Parallel()

	assert.Len(t, TextWithLen(100), 100)
}

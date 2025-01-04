package redact_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_RedactUtils_Redact_PasswordReturnRedacted(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "[REDACTED]", Redact(""))
	assert.Equal(t, "[REDACTED:3]", Redact("123"))
	assert.Equal(t, "[REDACTED:10]", Redact("1234567890"))
}

func TestUnit_RedactUtils_RedactWithPrefix_Ok(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "[REDACTED]", RedactWithPrefix(""))
	assert.Equal(t, "[REDACTED:3]", RedactWithPrefix("123"))
	assert.Equal(t, "[REDACTED:12345...:10]", RedactWithPrefix("1234567890"))
}

func TestUnit_RedactUtils_Trim_Ok(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "[]", Trim(""))
	assert.Equal(t, "[123]", Trim("123"))
	assert.Equal(t, "[12345...:14]", Trim("12345678901234"))
}

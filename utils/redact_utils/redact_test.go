package redact_utils

import (
	"testing"

	"github.com/teadove/teasutils/utils/json_utils"
	"github.com/teadove/teasutils/utils/test_utils"

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

func TestUnit_RedactUtils_RedactJSONWithPrefix_Ok(t *testing.T) {
	t.Parallel()

	ctx := test_utils.GetLoggedContext()

	values := map[string]any{
		"user": map[string]any{
			"name":     "TeaDove",
			"password": "123456789",
			"phone":    123456789,
		},
		"db": map[string]any{
			"host":     "localhost",
			"port":     "5432",
			"password": "123456789",
		},
	}

	assert.JSONEq(
		t,
		`{"db":{"host":"localhost","password":"[REDACTED:12345...:9]","port":"5432"},
"user":{"name":"TeaDove","password":"[REDACTED:12345...:9]","phone":"[REDACTED:12345...:9]"}}`,
		string(RedactJSONWithPrefix(ctx,
			json_utils.MarshalOrWarn(ctx, values),
			"user.password", "user.phone", "db.password",
		)),
	)
}

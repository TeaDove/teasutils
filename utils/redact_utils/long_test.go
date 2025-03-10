package redact_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teadove/teasutils/utils/json_utils"
	"github.com/teadove/teasutils/utils/test_utils"
)

func TestUnit_RedactUtils_ReductLongStrings_Ok(t *testing.T) {
	t.Parallel()

	ctx := test_utils.GetLoggedContext()

	values := map[string]any{
		"user": map[string]any{
			"name":     "TeaDove",
			"password": "1234567890123456789",
			"phone":    123456789,
		},
		"db": map[string]any{
			"host":     "localhost",
			"port":     "5432",
			"password": "1234567890123456789",
		},
	}

	assert.JSONEq(
		t,
		`{"db":{"host":"localhost","password":"[REDACTED:12345...:19]","port":"5432"},
"user":{"name":"TeaDove","password":"[REDACTED:12345...:19]","phone":123456789}}`,
		string(RedactLongStrings(json_utils.MarshalOrWarn(ctx, values))),
	)
}

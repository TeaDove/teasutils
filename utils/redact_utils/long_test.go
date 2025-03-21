package redact_utils

import (
	"encoding/json"
	"testing"

	"github.com/teadove/teasutils/utils/must_utils"

	"github.com/stretchr/testify/assert"
)

func TestUnit_RedactUtils_ReductLongStrings_Ok(t *testing.T) {
	t.Parallel()

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
		`{"db":{"host":"localhost","password":"[REDACTED:123456789012...:19]","port":"5432"},
"user":{"name":"TeaDove","password":"[REDACTED:123456789012...:19]","phone":123456789}}`,
		string(RedactLongStrings(must_utils.Must(json.Marshal(values)))),
	)
}

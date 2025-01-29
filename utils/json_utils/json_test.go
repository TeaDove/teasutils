package json_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teadove/teasutils/utils/test_utils"
)

func TestUnit_JsonUtils_MarshalOrRaw_Ok(t *testing.T) {
	t.Parallel()

	ctx := test_utils.GetLoggedContext()

	assert.JSONEq(t, `"abc"`, string(MarshalOrWarn(ctx, "abc")))
	assert.JSONEq(t, `{"a":"b"}`, string(MarshalOrWarn(ctx, map[string]string{"a": "b"})))

	assert.JSONEq(t,
		"{}",
		string(MarshalOrWarn(ctx, func() {})),
	)
}

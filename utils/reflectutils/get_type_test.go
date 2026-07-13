package reflectutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_ReflectUtils_GetFunctionName_Ok(t *testing.T) {
	t.Parallel()

	assert.Equal(
		t,
		"github.com/teadove/teasutils/utils/reflectutils.TestUnit_ReflectUtils_GetFunctionName_Ok func(*testing.T)",
		GetFunctionName(TestUnit_ReflectUtils_GetFunctionName_Ok),
	)
}

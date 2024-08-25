package refrect_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_ReflectUtils_GetTypesStringRepresentationIntPtr_Ok(t *testing.T) {
	t.Parallel()

	v := 1
	assert.Equal(t, "*int", GetTypesStringRepresentation(&v))
}

func TestUnit_ReflectUtils_GetTypesStringRepresentationInt_Ok(t *testing.T) {
	t.Parallel()

	v := 1
	assert.Equal(t, "int", GetTypesStringRepresentation(v))
}

func TestUnit_ReflectUtils_GetFunctionName_Ok(t *testing.T) {
	t.Parallel()

	assert.Equal(
		t,
		"github.com/teadove/teasutils/utils/refrect_utils.TestUnit_ReflectUtils_GetFunctionName_Ok",
		GetFunctionName(TestUnit_ReflectUtils_GetFunctionName_Ok),
	)
}

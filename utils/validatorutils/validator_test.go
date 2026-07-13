package validatorutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/teadove/teasutils/utils/validatorutils"
)

type sample struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Age   int    `validate:"gte=0,lte=130"`
}

func TestValidator_Valid(t *testing.T) {
	t.Parallel()

	require.NoError(t, validatorutils.Validator.Struct(sample{Name: "Tea", Email: "a@b.com", Age: 30}))
}

func TestValidator_Invalid(t *testing.T) {
	t.Parallel()

	err := validatorutils.Validator.Struct(sample{Name: "", Email: "not-an-email", Age: 200})
	assert.Error(t, err)
}

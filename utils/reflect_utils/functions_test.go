package reflect_utils

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_ReflectUtils_ConvertToWithCtxAndErr_Ok(t *testing.T) {
	t.Parallel()

	exp := "func(context.Context) error"

	assert.Equal(t, exp, fmt.Sprintf("%T", ConvertToWithCtxAndErr(func(_ context.Context) error { return nil })))
	assert.Equal(t, exp, fmt.Sprintf("%T", ConvertToWithCtxAndErr(func(_ context.Context) {})))
	assert.Equal(t, exp, fmt.Sprintf("%T", ConvertToWithCtxAndErr(func() error { return nil })))
	assert.Equal(t, exp, fmt.Sprintf("%T", ConvertToWithCtxAndErr(func() {})))
}

package mustutils

import (
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMust(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 42, Must(42, nil))
	assert.Panics(t, func() { Must(0, errors.New("boom")) })
}

func TestMustOk(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "ok", MustOk("ok", true))
	assert.Panics(t, func() { MustOk("", false) })
}

func TestMustNoReturn(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() { MustNoReturn(nil) })
	assert.Panics(t, func() { MustNoReturn(errors.New("boom")) })
}

func TestAnyToErr(t *testing.T) {
	t.Parallel()

	require.NoError(t, AnyToErr(nil))
	require.ErrorContains(t, AnyToErr("oops"), "oops")
	require.ErrorContains(t, AnyToErr(errors.New("boom")), "boom")
	require.ErrorContains(t, AnyToErr(42), "42")
}

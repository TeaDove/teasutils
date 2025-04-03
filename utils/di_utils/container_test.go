package di_utils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestUnit_DIUtils_MustBuildFromSettingsAndRun_Ok(t *testing.T) {
	t.Parallel()

	called := false

	require.NoError(t, checkFromCheckers(context.Background(), []any{func() { called = true }}))
	assert.True(t, called)
}

func TestUnit_DIUtils_MustBuildFromSettingsStopInf_LogsErr(t *testing.T) {
	t.Parallel()

	stoped := false
	err := stop(context.Background(), []any{func() { stoped = true }})

	require.NoError(t, err)
	assert.True(t, stoped)
}

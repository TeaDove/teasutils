package concutils

import (
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/teadove/teasutils/utils/randomutils"
)

func TestMapGroup(t *testing.T) {
	t.Parallel()

	group := NewMapGroup[uuid.UUID, string](10, 2)
	for range 10 {
		group.Go(func() (uuid.UUID, string, error) {
			return uuid.New(), randomutils.TextWithLen(10), nil
		})
	}

	mapping, err := group.Collect()
	require.NoError(t, err)
	assert.Len(t, mapping, 10)
}

// TestMapGroup_ConcurrentErrors joins every error and keeps successful
// entries. Run with -race to catch unsynchronised writes to the error/map.
func TestMapGroup_ConcurrentErrors(t *testing.T) {
	t.Parallel()

	const total = 50

	group := NewMapGroup[int, int](total, 8)
	for i := range total {
		group.Go(func() (int, int, error) {
			if i%2 == 0 {
				return 0, 0, errors.Newf("fail-%d", i)
			}

			return i, i * i, nil
		})
	}

	mapping, err := group.Collect()
	require.Error(t, err)
	assert.Len(t, mapping, total/2, "odd keys must survive")

	for k, v := range mapping {
		assert.Equal(t, k*k, v)
	}
}

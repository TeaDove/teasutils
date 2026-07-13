package maputils

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teadove/teasutils/utils/randomutils"
)

func TestRMMap(t *testing.T) {
	t.Parallel()

	rwMap := NewRWMap[string, string](1000)

	var wg sync.WaitGroup
	for range 1000 {
		wg.Go(func() {
			rwMap.Set(randomutils.TextWithLen(6), randomutils.TextWithLen(10))
		})

		wg.Go(func() {
			name, ok := rwMap.Get(randomutils.TextWithLen(6))
			if ok {
				assert.NotEmpty(t, name)
			}
		})
	}

	wg.Wait()
	assert.GreaterOrEqual(t, len(rwMap.Copy()), 100)

	for key, value := range rwMap.Copy() {
		assert.NotEmpty(t, key)
		assert.NotEmpty(t, value)
	}
}

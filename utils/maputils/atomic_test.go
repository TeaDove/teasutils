package maputils

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func genValue() (string, string) {
	id := rand.IntN(size)

	return fmt.Sprintf("key-%06d", id), fmt.Sprintf("val-%06d", id)
}

const size = 1_000

func newMap() map[string]string {
	v := make(map[string]string, size)

	for {
		key, value := genValue()
		v[key] = value

		if len(v) >= size {
			return v
		}
	}
}

func TestAtomicMap(t *testing.T) {
	t.Parallel()

	atomicMap := AtomicMap[string, string]{}
	atomicMap.Store(newMap())
	assert.Equal(t, size, atomicMap.Len())

	var (
		needToFind atomic.Int64
		wg         sync.WaitGroup
	)
	needToFind.Store(size / 100)

	for range 100 {
		wg.Go(func() {
			for {
				k, _ := genValue()

				realV, ok := atomicMap.Get(k)
				if ok {
					assert.NotEmpty(t, k)
					assert.NotEmpty(t, realV)
					// fmt.Printf("%s: %s\n", k, realV)

					loaded := needToFind.Add(-1)
					if loaded <= 0 {
						return
					}

					if loaded%50 == 0 {
						atomicMap.Store(newMap())
					}
				}

				time.Sleep(100 * time.Millisecond)
			}
		})
	}

	wg.Wait()
}

func TestNewAtomicMapWithRefresher(t *testing.T) {
	t.Parallel()

	atomicMap, err := NewRefreshingAtomicMap(t.Context(), time.Second, func() (map[string]string, error) {
		return newMap(), nil
	})
	require.NoError(t, err)
	assert.Equal(t, size, atomicMap.atomicMap.Len())

	err = atomicMap.Refresh(t.Context())
	require.NoError(t, err)
	assert.Equal(t, size, atomicMap.atomicMap.Len())
}

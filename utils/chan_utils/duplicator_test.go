package chan_utils

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestUnit_ChanUtils_Duplicator_Ok(t *testing.T) {
	var wg sync.WaitGroup

	original := make(chan string, 10)
	a, b := Duplicator(original, &wg)

	original <- "a"
	original <- "b"
	original <- "c"

	close(original)
	wg.Wait()
	close(a)
	close(b)

	assert.Equal(t, []string{"a", "b", "c"}, ChanToSlice(a))
	assert.Equal(t, []string{"a", "b", "c"}, ChanToSlice(b))

	wg.Wait()
}

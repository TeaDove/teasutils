package test_utils

import "testing"

func SkipIfShortMode(t *testing.T) {
	t.Helper()

	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
}

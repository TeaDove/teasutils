package mustutils

import (
	"context"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/teadove/teasutils/utils/testutils"
)

func panicingFunction(_ context.Context) error {
	panic("bad panic")
}

func failingFunction(_ context.Context) error {
	return errors.New("bad error")
}

func TestDoOrLog(t *testing.T) {
	t.Parallel()

	WithRecoverAndLog(panicingFunction)(testutils.Context())
	WithRecoverAndLog(failingFunction)(testutils.Context())
}

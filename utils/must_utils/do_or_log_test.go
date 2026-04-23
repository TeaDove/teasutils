package must_utils

import (
	"context"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/teadove/teasutils/utils/test_utils"
)

func panicingFunction(_ context.Context) error {
	panic("bad panic")
}

func failingFunction(_ context.Context) error {
	return errors.New("bad error")
}

func TestDoOrLog(t *testing.T) {
	t.Parallel()

	WithRecoverAndLog(panicingFunction)(test_utils.GetLoggedContext())
	WithRecoverAndLog(failingFunction)(test_utils.GetLoggedContext())
}

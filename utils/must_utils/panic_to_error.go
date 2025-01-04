package must_utils

import (
	"fmt"

	"github.com/pkg/errors"
)

func AnyToErr(v any) error {
	if v == nil {
		return nil
	}

	err, ok := v.(error)
	if ok {
		return errors.Wrap(err, "paniced")
	}

	strErr, ok := v.(string)
	if ok {
		return errors.Wrap(errors.New("paniced"), strErr)
	}

	return fmt.Errorf("paniced: %v", v)
}

package reflect_utils

import "context"

// ConvertToWithCtxAndErr
// Convert functions with sig. like func(ctx context.Context) to `func(ctx context.Context) error`.
func ConvertToWithCtxAndErr(callable any) func(ctx context.Context) error {
	switch call := callable.(type) {
	case func():
		return func(_ context.Context) error {
			call()
			return nil
		}
	case func() error:
		return func(_ context.Context) error {
			return call()
		}
	case func(context.Context):
		return func(ctx context.Context) error {
			call(ctx)
			return nil
		}
	case func(ctx context.Context) error:
		return call
	default:
		return nil
	}
}

package must_utils

func Must[T any](obj T, err error) T {
	if err != nil {
		panic("must failed")
	}
	return obj
}

func MustNoReturn(err error) {
	if err != nil {
		panic("must failed")
	}
}

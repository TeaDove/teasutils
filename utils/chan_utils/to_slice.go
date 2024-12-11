package chan_utils

func ChanToSlice[T any](channel <-chan T) []T {
	slice := make([]T, 0)
	var v T
	for v = range channel {
		slice = append(slice, v)
	}

	return slice
}

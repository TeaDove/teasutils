package slices_utils

func PythoniseIdx[T any](slice []T, idx int) int {
	if idx >= 0 {
		return idx
	}

	return len(slice) + idx
}

func PythoniseIdxGet[T any](slice []T, idx int) T {
	if idx >= 0 {
		return slice[idx]
	}

	return slice[len(slice)+idx]
}

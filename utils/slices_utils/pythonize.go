package slices_utils

func PythoniseIdx[T ~[]E, E any](slice T, idx int) int {
	if idx >= 0 {
		return idx
	}

	return len(slice) + idx
}

func PythoniseIdxGet[T ~[]E, E any](slice T, idx int) E {
	if idx >= 0 {
		return slice[idx]
	}

	return slice[len(slice)+idx]
}

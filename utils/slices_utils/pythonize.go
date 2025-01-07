package slices_utils

func PIdx[T []E, E any](slice T, idx int) int {
	if idx >= 0 {
		return idx
	}

	return len(slice) + idx
}

func PSIdx(slice string, idx int) int {
	if idx >= 0 {
		return idx
	}

	return len(slice) + idx
}

func PGet[T []E, E any](slice T, idx int) E {
	if idx >= 0 {
		return slice[idx]
	}

	return slice[len(slice)+idx]
}

func PSGet(slice string, idx int) byte {
	if idx >= 0 {
		return slice[idx]
	}

	return slice[len(slice)+idx]
}

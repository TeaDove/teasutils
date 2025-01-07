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
	return slice[PIdx(slice, idx)]
}

func PAGet[T []E, E any](slice T, start int, end int) []E {
	if end >= 0 {
		if start >= 0 {
			return slice[start:end]
		}

		return slice[len(slice)-start : end]
	}

	return slice[PIdx(slice, start):PIdx(slice, end)]
}

func PSGet(slice string, idx int) byte {
	if idx >= 0 {
		return slice[idx]
	}

	return slice[PSIdx(slice, idx)]
}

func PSAGet(slice string, start int, end int) string {
	return slice[PSIdx(slice, start):PSIdx(slice, end)]
}

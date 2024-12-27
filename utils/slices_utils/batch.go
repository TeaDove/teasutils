package slices_utils

// CutInBatches
// Cuts slice in batches, remaining will be put in extra batch.
func CutInBatches[T any](slice []T, batchSize int) [][]T {
	if batchSize <= 0 {
		panic("batch size must be greater than zero")
	}

	if len(slice) == 0 {
		return [][]T{}
	}

	if batchSize >= len(slice) {
		return [][]T{slice}
	}

	batches := make([][]T, 0, len(slice)/batchSize)

	var i int
	for i = range len(slice) / batchSize {
		batches = append(batches, slice[i*batchSize:(i+1)*batchSize])
	}

	i++

	if (i)*batchSize < len(slice) {
		batches = append(batches, slice[(i)*batchSize:])
	}

	return batches
}

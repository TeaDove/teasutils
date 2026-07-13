package mathutils

import "golang.org/x/exp/constraints"

type number interface {
	constraints.Integer | constraints.Float
}

// Avg returns the arithmetic mean of xs. It returns NaN for an empty slice
// (division by zero), so the caller must ensure len(xs) > 0.
func Avg[T number](xs []T) float64 {
	var sum T
	for _, x := range xs {
		sum += x
	}

	return float64(sum) / float64(len(xs))
}

// AddToAvg returns the running average after adding x to a set of n values
// whose current mean is originalAvg.
func AddToAvg[T number](originalAvg float64, n int, x T) float64 {
	return (originalAvg*float64(n) + float64(x)) / float64(n+1)
}

// AvgWithAvg merges two averages (avgA over na items, avgb over nb items)
// into the combined mean. It returns NaN when na+nb == 0.
func AvgWithAvg(avgA float64, na int, avgb float64, nb int) float64 {
	return (avgA*float64(na) + avgb*float64(nb)) / float64(na+nb)
}

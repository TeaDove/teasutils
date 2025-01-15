package math_utils

import "golang.org/x/exp/constraints"

type number interface {
	constraints.Integer | constraints.Float
}

func Avg[T number](xs []T) float64 {
	var sum T
	for _, x := range xs {
		sum += x
	}

	return float64(sum) / float64(len(xs))
}

func AddToAvg[T number](originalAvg float64, n int, x T) float64 {
	return (originalAvg*float64(n) + float64(x)) / float64(n+1)
}

func AvgWithAvg(avgA float64, na int, avgb float64, nb int) float64 {
	return (avgA*float64(na) + avgb*float64(nb)) / float64(na+nb)
}

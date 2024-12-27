package converters_utils

import (
	"math"

	"golang.org/x/exp/constraints"
)

const (
	thousand       = 1000
	binaryThousand = 1024
	precision      = 2
)

func round(num float64) int {
	//nolint: mnd // its just 0.5...
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	//nolint: mnd // its just 10...
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func ToKilo[T constraints.Integer](bytes T) float64 {
	return ToFixed(float64(bytes)/thousand, precision)
}

func ToKiloByte[T constraints.Integer](bytes T) float64 {
	return ToFixed(float64(bytes)/binaryThousand, precision)
}

func ToMega[T constraints.Integer](bytes T) float64 {
	return ToFixed(float64(bytes)/thousand/thousand, precision)
}

func ToMegaByte[T constraints.Integer](bytes T) float64 {
	return ToFixed(float64(bytes)/binaryThousand/binaryThousand, precision)
}

func ToGiga[T constraints.Integer](bytes T) float64 {
	return ToFixed(float64(bytes)/thousand/thousand/thousand, precision)
}

func ToGigaByte[T constraints.Integer](bytes T) float64 {
	return ToFixed(float64(bytes)/binaryThousand/binaryThousand/binaryThousand, precision)
}

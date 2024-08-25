package converters_utils

import (
	"math"

	"golang.org/x/exp/constraints"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func ToKilo[T constraints.Integer](bytes T) float64 {
	return ToFixed(float64(bytes)/1000, 2)
}

func ToKiloByte[T constraints.Integer](bytes T) float64 {
	return ToFixed(float64(bytes)/1024, 2)
}

func ToMega[T constraints.Integer](bytes T) float64 {
	return ToFixed(float64(bytes)/1000/1000, 2)
}

func ToMegaByte[T constraints.Integer](bytes T) float64 {
	return ToFixed(float64(bytes)/1024/1024, 2)
}

func ToGiga[T constraints.Integer](bytes T) float64 {
	return ToFixed(float64(bytes)/1000/1000/1000, 2)
}

func ToGigaByte[T constraints.Integer](bytes T) float64 {
	return ToFixed(float64(bytes)/1024/1024/1024, 2)
}

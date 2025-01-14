package conv_utils

import (
	"fmt"
	"math"
	"strconv"

	"golang.org/x/exp/constraints"
)

const (
	thousand         = 1000
	binaryThousand   = 1024
	defaultPrecision = 2
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

func ToKilo[T constraints.Integer | constraints.Float](bytes T) float64 {
	return float64(bytes) / thousand
}

func ToKiloByte[T constraints.Integer | constraints.Float](bytes T) float64 {
	return float64(bytes) / binaryThousand
}

func ToMega[T constraints.Integer | constraints.Float](bytes T) float64 {
	return float64(bytes) / thousand / thousand
}

func ToMegaByte[T constraints.Integer | constraints.Float](bytes T) float64 {
	return float64(bytes) / binaryThousand / binaryThousand
}

func ToGiga[T constraints.Integer | constraints.Float](bytes T) float64 {
	return float64(bytes) / thousand / thousand / thousand
}

func ToGigaByte[T constraints.Integer | constraints.Float](bytes T) float64 {
	return float64(bytes) / binaryThousand / binaryThousand / binaryThousand
}

func ClosestByteAndK[T constraints.Integer | constraints.Float](bytes T) (float64, uint) {
	if float64(bytes) <= binaryThousand {
		return float64(bytes), 0 // B
	}

	if float64(bytes) <= binaryThousand*binaryThousand {
		return float64(bytes) / binaryThousand, 1 // KB
	}

	if float64(bytes) <= binaryThousand*binaryThousand*binaryThousand {
		//nolint: mnd // go fuck yourself
		return float64(bytes) / binaryThousand / binaryThousand, 2 // MB
	}

	if float64(bytes) <= binaryThousand*binaryThousand*binaryThousand*binaryThousand {
		//nolint: mnd // go fuck yourself
		return float64(bytes) / binaryThousand / binaryThousand / binaryThousand, 3 // GB
	}

	//nolint: mnd // go fuck yourself
	return float64(bytes) / binaryThousand / binaryThousand / binaryThousand / binaryThousand, 4 // TB
}

func ClosestByteWithPrecision[T constraints.Integer | constraints.Float](bytes T, precision int) string {
	rounded, pow := ClosestByteAndK(bytes)

	var digit string

	switch pow {
	case 0:
		digit = "B"

	case 1:
		digit = "KB"
	//nolint: mnd // go fuck yourself
	case 2:
		digit = "MB"
	//nolint: mnd // go fuck yourself
	case 3:
		digit = "GB"
	default:
		digit = "TB"
	}

	return fmt.Sprintf(
		"%s %s",
		strconv.FormatFloat(ToFixed(rounded, precision), 'f', -1, 64),
		digit,
	)
}

func ClosestByte[T constraints.Integer | constraints.Float](bytes T) string {
	return ClosestByteWithPrecision(bytes, defaultPrecision)
}

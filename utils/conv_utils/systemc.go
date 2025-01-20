package conv_utils

import (
	"fmt"
	"math"
	"strconv"

	"golang.org/x/exp/constraints"
)

const (
	thousand         = 1000.0
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

func closestByteAndK(v float64) (float64, uint) {
	if v <= binaryThousand {
		return v, 0 // B
	}

	if v <= binaryThousand*binaryThousand {
		return v / binaryThousand, 1 // KB
	}

	if v <= binaryThousand*binaryThousand*binaryThousand {
		//nolint: mnd // go fuck yourself
		return v / binaryThousand / binaryThousand, 2 // MB
	}

	if v <= binaryThousand*binaryThousand*binaryThousand*binaryThousand {
		//nolint: mnd // go fuck yourself
		return v / binaryThousand / binaryThousand / binaryThousand, 3 // GB
	}

	//nolint: mnd // go fuck yourself
	return v / binaryThousand / binaryThousand / binaryThousand / binaryThousand, 4 // TB
}

func ClosestByteWithPrecision[T constraints.Integer | constraints.Float](bytes T, precision int) string {
	rounded, pow := closestByteAndK(float64(bytes))

	var digit string

	switch pow {
	case 0:
		digit = "B"

	case 1:
		digit = "kB"
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

func closestK(v float64) (float64, int) {
	if v <= (1 / (thousand * thousand)) {
		return v * thousand * thousand * thousand, -3 // Nano
	}

	if v <= (1 / thousand) {
		return v * thousand * thousand, -2 // Micro
	}

	if v <= 1 {
		return v * thousand, -1 // Milli
	}

	if v <= thousand {
		return v, 0 // B
	}

	if v <= thousand*thousand {
		return v / thousand, 1 // KB
	}

	if v <= thousand*thousand*thousand {
		//nolint: mnd // go fuck yourself
		return v / thousand / thousand, 2 // MB
	}

	if v <= thousand*thousand*thousand*thousand {
		//nolint: mnd // go fuck yourself
		return v / thousand / thousand / thousand, 3 // GB
	}

	//nolint: mnd // go fuck yourself
	return v / thousand / thousand / thousand / thousand, 4 // TB
}

func ClosestKWithPrecision[T constraints.Integer | constraints.Float](v T, precision int) string {
	rounded, pow := closestK(float64(v))

	var digit string

	switch pow {
	case -3:
		digit = " n"
	case -2:
		digit = " µ"
	case -1:
		digit = " m"
	case 0:
		digit = ""

	case 1:
		digit = " k"
	//nolint: mnd // go fuck yourself
	case 2:
		digit = " M"
	//nolint: mnd // go fuck yourself
	case 3:
		digit = " G"
	default:
		digit = " T"
	}

	return fmt.Sprintf(
		"%s%s",
		strconv.FormatFloat(ToFixed(rounded, precision), 'f', -1, 64),
		digit,
	)
}

func Closest[T constraints.Integer | constraints.Float](v T) string {
	return ClosestKWithPrecision(v, defaultPrecision)
}

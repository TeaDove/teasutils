package convutils

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

// ToFixed rounds num to precision decimal places.
func ToFixed(num float64, precision int) float64 {
	//nolint: mnd // its just 10...
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

// ToKilo divides by 1000 (decimal kilo).
func ToKilo[T constraints.Integer | constraints.Float](bytes T) float64 {
	return float64(bytes) / thousand
}

// ToKiloByte divides by 1024 (binary kibi).
func ToKiloByte[T constraints.Integer | constraints.Float](bytes T) float64 {
	return float64(bytes) / binaryThousand
}

// ToMega divides by 1000^2 (decimal mega).
func ToMega[T constraints.Integer | constraints.Float](bytes T) float64 {
	return float64(bytes) / thousand / thousand
}

// ToMegaByte divides by 1024^2 (binary mebi).
func ToMegaByte[T constraints.Integer | constraints.Float](bytes T) float64 {
	return float64(bytes) / binaryThousand / binaryThousand
}

// ToGiga divides by 1000^3 (decimal giga).
func ToGiga[T constraints.Integer | constraints.Float](bytes T) float64 {
	return float64(bytes) / thousand / thousand / thousand
}

// ToGigaByte divides by 1024^3 (binary gibi).
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

// ClosestByteWithPrecision formats bytes with the largest fitting binary unit
// (B, kB, MB, GB, TB) and precision decimals, e.g. "1.5 MB". Intended for
// non-negative sizes.
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

// ClosestByte is ClosestByteWithPrecision with the default precision (2).
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

// ClosestKWithPrecision formats v with the closest decimal SI prefix, from
// nano ("n") up to tera ("T"), with precision decimals, e.g. "1.5 k".
// Intended for non-negative magnitudes.
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

// Closest is ClosestKWithPrecision with the default precision (2).
func Closest[T constraints.Integer | constraints.Float](v T) string {
	return ClosestKWithPrecision(v, defaultPrecision)
}

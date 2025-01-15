package math_utils

import "github.com/teadove/teasutils/utils/conv_utils"

const slicePrecision = 3

func SliceLinear(target float64, count int) []float64 {
	var (
		step = target / float64(count)
		base = step
		res  = make([]float64, 0, count)
	)

	for range count {
		res = append(res, conv_utils.ToFixed(base, slicePrecision))
		base += step
	}

	return res
}

func SliceGeometricProgression(base float64, commonRation float64, count int) []float64 {
	var (
		v   = base
		res = make([]float64, 0, count)
	)

	res = append(res, v)

	for range count {
		v *= commonRation
		res = append(res, conv_utils.ToFixed(v, slicePrecision))
	}

	return res
}

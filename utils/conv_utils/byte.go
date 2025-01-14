package conv_utils

const defaultPrecision = 2

type Byte struct {
	v float64
}

func (r *Byte) String() string {
	return ToClosestByteAsString(r.v, defaultPrecision)
}

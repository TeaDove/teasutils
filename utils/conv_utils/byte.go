package conv_utils

const defaultPrecision = 2

type Byte float64

func (r Byte) String() string {
	return ToClosestByteAsString(r, defaultPrecision)
}

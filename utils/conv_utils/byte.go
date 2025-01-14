package conv_utils

type Byte float64

func (r Byte) String() string {
	return ClosestByte(r)
}

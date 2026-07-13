package convutils

// Byte is a byte count that renders itself with a human-readable binary unit.
type Byte float64

// String formats the value via ClosestByte, e.g. "1.5 MB".
func (r Byte) String() string {
	return ClosestByte(r)
}

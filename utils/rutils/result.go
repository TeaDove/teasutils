package rutils

// Result bundles the value and error returned by a fallible operation,
// so it can be passed over a channel or stored in a slice.
type Result[T any] struct {
	Ok  T     `json:"ok,omitempty"`
	Err error `json:"err,omitempty"`
}

// NewResult wraps an (value, error) pair into a Result.
func NewResult[T any](ok T, err error) Result[T] {
	return Result[T]{Ok: ok, Err: err}
}

// Pair holds two related values of independent types.
type Pair[T, K any] struct {
	First  T `json:"first,omitempty"`
	Second K `json:"second,omitempty"`
}

// NewPair builds a Pair from its two components.
func NewPair[T, K any](first T, second K) Pair[T, K] {
	return Pair[T, K]{First: first, Second: second}
}

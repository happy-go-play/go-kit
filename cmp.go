package gokit

import "cmp"

// numeric is the constraint for types that support negation and comparison with zero (signed integers and floats).
type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

// Abs returns the absolute value of x.
// T must be a signed integer or float type.
func Abs[T numeric](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Max returns the larger of a and b.
// T must be an ordered type (supports <, <=, >, >=).
func Max[T cmp.Ordered](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

// Min returns the smaller of a and b.
// T must be an ordered type (supports <, <=, >, >=).
func Min[T cmp.Ordered](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

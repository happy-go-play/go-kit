package gokit

func FromPtr[T any](s *T) T {
	if s == nil {
		var zero T
		return zero
	}
	return *s
}

// Deprecated: Use go1.26's new(T) instead
func ToPtr[T any](s T) *T {
	return &s
}

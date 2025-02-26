package gokit

func FromPtr[T any](s *T) T {
	if s == nil {
		var zero T
		return zero
	}
	return *s
}

func ToPtr[T any](s T) *T {
	return &s
}

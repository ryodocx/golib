package ptr

func ToPtr[T any](input T) *T {
	return &input
}

func FromPtr[T any](ptr *T, defaultValue T) T {
	if ptr != nil {
		return *ptr
	}
	return defaultValue
}

package pdf

func ResizeSlice[T any](slice []T, size int) []T {
	if cap(slice) >= size {
		return slice[:size]
	}

	newSlice := make([]T, size)
	copy(newSlice, slice)

	return newSlice
}

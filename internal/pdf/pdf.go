package pdf

func ResizeSlice[T any](slice []T, size int) []T {
	if cap(slice) >= size {
		return slice[:size]
	}

	newSlice := make([]T, size)
	copy(newSlice, slice)

	return newSlice
}

func CheckEOL(e1, e2 byte) bool {
	// From pdf reference, page 94:
	// If the file's end-of-line marker is a single character (either a carriage return or a line feed),
	// it is preceded by a single space; if the marker is 2 characters (both a carriage return and a line feed),
	// it is not preceded by a space.
	return ((e1 == '\r' && e2 == '\n') ||
		(e1 == '\n' && e2 == '\r') ||
		(e1 == ' ' && (e2 == '\r' || e2 == '\n')))
}

func CheckXRefEntryType(c byte) bool {
	return c == 'n' || c == 'f'
}

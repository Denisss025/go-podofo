package pdf

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func ResizeSlice[T any](slice []T, size int) []T {
	if cap(slice) >= size {
		return slice[:size]
	}

	newSlice := make([]T, size)
	copy(newSlice, slice)

	return newSlice
}

func CheckEOL(e1, e2 byte) bool {
	const (
		crlf uint16 = 0x0a0d
		lfcr uint16 = 0x0d0a
		splf uint16 = 0x200a
		spcr uint16 = 0x200d
	)

	e := (uint16(e1) << 8) | uint16(e2)

	// From pdf reference, page 94:
	// If the file's end-of-line marker is a single character (either a carriage return or a line feed),
	// it is preceded by a single space; if the marker is 2 characters (both a carriage return and a line feed),
	// it is not preceded by a space.
	return e == crlf || e == lfcr || e == splf || e == spcr
}

func CheckXRefEntryType(c byte) bool {
	return c == 'n' || c == 'f'
}

func IsWhitespace(r rune) bool {
	const whiteSpaces = "\t\n\f\r "

	return r == 0 || strings.ContainsRune(whiteSpaces, r)
}

func IndexLast(r io.ReadSeeker, what []byte) (pos int64, err error) {
	buf := make([]byte, len(what))
	size, err := r.Seek(0, io.SeekEnd)
	if err != nil {
		return -1, fmt.Errorf("last index: %w", err)
	}

	for pos = size - int64(len(buf)); pos >= 0; pos-- {
		_, err = r.Seek(pos, io.SeekStart)
		if err != nil {
			return -1, fmt.Errorf("last index: %w", err)
		}

		_, err = r.Read(buf)
		if err != nil {
			return -1, fmt.Errorf("last index: %w", err)
		}

		if bytes.Equal(buf, what) {
			return pos, nil
		}
	}

	return -1, nil
}

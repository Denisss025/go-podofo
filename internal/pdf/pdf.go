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
	const bufSize = 4096

	savePos, _ := r.Seek(0, io.SeekCurrent)

	pos, err = r.Seek(0, io.SeekEnd)
	if err != nil {
		return -1, fmt.Errorf("index last: %w", err)
	}

	var buf []byte

	if pos > bufSize {
		buf = make([]byte, bufSize)
	} else {
		buf = make([]byte, pos)
	}

	_, err = r.Seek(-int64(len(buf)), io.SeekEnd)
	if err != nil {
		return -1, fmt.Errorf("last index: %w", err)
	}

	for {
		n, err := r.Read(buf)
		if err != nil {
			return -1, fmt.Errorf("last index: %w", err)
		}

		pos -= int64(n)

		i := bytes.LastIndex(buf, what)
		if i >= 0 {
			pos += int64(i)

			break
		}

		if pos == 0 {
			pos--

			break
		}

		if pos < bufSize {
			buf = buf[:pos]
		}
	}

	if _, err = r.Seek(savePos, io.SeekStart); err != nil {
		err = fmt.Errorf("index last: %w", err)
	}

	return pos, err
}

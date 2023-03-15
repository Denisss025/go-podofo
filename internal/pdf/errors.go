package pdf

import "errors"

var (
	ErrValueOutOfRange = errors.New("value out of range")

	ErrNotImplemented = errors.New("not implemented")

	ErrCannotConvertColor = errors.New("cannot convert color")
)

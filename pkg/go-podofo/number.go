package podofo

import "io"

type Number string

func (num Number) Kind() ObjectKind { return ObjectKindNumber }

func (num Number) WriteTo(w io.Writer) (n int64, err error) {
	panic("not implemented") // TODO: implement me
}

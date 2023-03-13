package podofo

import "io"

type Array struct {
	Objects []Object
}

func (array *Array) Kind() ObjectKind { return ObjectKindArray }

func (array *Array) WriteTo(w io.Writer) (n int64, err error) {
	panic("not implemented") // TODO: implement me
}

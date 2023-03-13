package podofo

import "io"

type Bool bool

func (b Bool) Kind() ObjectKind { return ObjectKindBool }

func (b Bool) WriteTo(w io.Writer) (n int64, err error) {
	panic("not implemented") // TODO: implement me
}

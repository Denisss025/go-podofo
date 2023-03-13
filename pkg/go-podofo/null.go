package podofo

import "io"

type Null struct{}

func (null Null) Kind() ObjectKind { return ObjectKindNull }

func (null Null) WriteTo(w io.Writer) (n int64, err error) {
	panic("not implemented") // TODO: implement me
}

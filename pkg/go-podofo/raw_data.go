package podofo

import "io"

type RawData struct{}

func (data *RawData) Kind() ObjectKind { return ObjectKindRawData }

func (data *RawData) WriteTo(w io.Writer) (n int64, err error) {
	panic("not implemented") // TODO: implement me
}

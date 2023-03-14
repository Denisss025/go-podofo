package podofo

import (
	"io"

	"github.com/denisss025/go-podofo/internal/pdf"
)

type Bool bool

func (b Bool) Kind() ObjectKind { return pdf.ObjectKindBool }

func (b Bool) WriteTo(w io.Writer) (n int64, err error) {
	panic("not implemented") // TODO: implement me
}

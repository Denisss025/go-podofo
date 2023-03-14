package podofo

import (
	"io"

	"github.com/denisss025/go-podofo/internal/pdf"
)

type Null struct{}

func (null Null) Kind() ObjectKind { return pdf.ObjectKindNull }

func (null Null) WriteTo(w io.Writer) (n int64, err error) {
	panic("not implemented") // TODO: implement me
}

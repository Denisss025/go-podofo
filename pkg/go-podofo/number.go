package podofo

import (
	"io"

	"github.com/denisss025/go-podofo/internal/pdf"
)

type Number string

func (num Number) Kind() ObjectKind { return pdf.ObjectKindNumber }

func (num Number) WriteTo(w io.Writer) (n int64, err error) {
	panic("not implemented") // TODO: implement me
}

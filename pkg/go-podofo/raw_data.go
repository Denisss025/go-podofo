package podofo

import (
	"io"

	"github.com/denisss025/go-podofo/internal/pdf"
)

type RawData struct{}

func (data *RawData) Kind() ObjectKind { return pdf.ObjectKindRawData }

func (data *RawData) WriteTo(w io.Writer) (n int64, err error) {
	panic("not implemented") // TODO: implement me
}

package podofo

import (
	"github.com/denisss025/go-podofo/internal/pdf"
)

type Name string

// TODO: need to make Name an object

const (
	KeyNull     Name = ""
	KeyContents Name = "Contents"
	KeyEncrypt  Name = "Encrypt"
	KeyFilter   Name = "Filter"
	KeyFlags    Name = "Flags"
	KeyID       Name = "ID"
	KeyInfo     Name = "Info"
	KeyLength   Name = "Length"
	KeyRect     Name = "Rect"
	KeyRoot     Name = "Root"
	KeySize     Name = "Size"
	KeySubtype  Name = "Subtype"
	KeyType     Name = "Type"

	NameXRef Name = "XRef"
)

func (name Name) Kind() ObjectKind { return pdf.ObjectKindName }

func (name Name) MarshalPDF(w *pdf.Writer) error {
	panic("not implemented") // TODO: implement me
}

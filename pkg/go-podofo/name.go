package podofo

import "io"

type Name string

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

func (name Name) Kind() ObjectKind { return ObjectKindName }

func (name Name) WriteTo(w io.Writer) (n int64, err error) {
	panic("not implemented") // TODO: implement me
}

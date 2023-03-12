package podofo

type Name string

const (
	NameKeyNull     Name = ""
	NameKeyContents Name = "Contents"
	NameKeyEncrypt  Name = "Encrypt"
	NameKeyFilter   Name = "Filter"
	NameKeyFlags    Name = "Flags"
	NameKeyID       Name = "ID"
	NameKeyLength   Name = "Length"
	NameKeyRect     Name = "Rect"
	NameKeySize     Name = "Size"
	NameKeySubtype  Name = "Subtype"
	NameKeyType     Name = "Type"

	NameXRef Name = "XRef"
)

func (name Name) Kind() ObjectKind { return ObjectKindName }

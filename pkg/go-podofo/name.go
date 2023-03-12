package podofo

type Name string

const (
	NameKeyNull     Name = ""
	NameKeyContents Name = "Contents"
	NameKeyFlags    Name = "Flags"
	NameKeyLength   Name = "Length"
	NameKeyRect     Name = "Rect"
	NameKeySize     Name = "Size"
	NameKeySubtype  Name = "Subtype"
	NameKeyType     Name = "Type"
	NameKeyFilter   Name = "Filter"
	NameKeyEncrypt  Name = "Encrypt"

	NameXRef Name = "XRef"
)

func (name Name) Kind() ObjectKind { return ObjectKindName }

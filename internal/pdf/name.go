package pdf

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

type NameObject struct {
	BaseObject

	Name Name
}

func (name *NameObject) Kind() ObjectKind { return ObjectKindName }

func (name *NameObject) MarshalPDF(w *Writer) error {
	panic("not implemented") // TODO: implement me
}

func (name *NameObject) SetParent(parent Object) {
	panic("not implemented") // TODO: Implement
}

func (name *NameObject) Parent() Object {
	panic("not implemented") // TODO: Implement
}

func (name *NameObject) Copy() (Object, error) {
	panic("not implemented") // TODO: Implement
}

func (name *NameObject) Document() *Document {
	panic("not implemented") // TODO: Implement
}

func (name *NameObject) GetIndirectReference() *Reference {
	panic("not implemented") // TODO: Implement
}

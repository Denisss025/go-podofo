package pdf

type Dictionary struct{}

// MarshalPDF encodes the receiver a PDF bytes.
func (dict *Dictionary) MarshalPDF(_ *Writer) error {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) SetParent(_ Object) {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) Parent() Object {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) Kind() ObjectKind {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) Copy() (Object, error) {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) Document() *Document {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) GetIndirectReference() *Reference {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) Key(name Name) Object {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) AddKeyIndirect(name Name, obj Object) {
	panic("not implemented") // TODO: Implement
}

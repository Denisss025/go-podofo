package pdf

type Array struct{}

// MarshalPDF encodes the receiver a PDF bytes.
func (arr *Array) MarshalPDF(_ *Writer) error {
	panic("not implemented") // TODO: Implement
}

func (arr *Array) SetParent(_ Object) {
	panic("not implemented") // TODO: Implement
}

func (arr *Array) Parent() Object {
	panic("not implemented") // TODO: Implement
}

func (arr *Array) Kind() ObjectKind {
	panic("not implemented") // TODO: Implement
}

func (arr *Array) Copy() (Object, error) {
	panic("not implemented") // TODO: Implement
}

func (arr *Array) Document() *Document {
	panic("not implemented") // TODO: Implement
}

func (arr *Array) GetIndirectReference() *Reference {
	panic("not implemented") // TODO: Implement
}

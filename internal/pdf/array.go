package pdf

type Array struct{}

// MarshalPDF encodes the receiver a PDF bytes.
func (arr *Array) MarshalPDF(_ *pdf.Writer) error {
	panic("not implemented") // TODO: Implement
}

func (arr *Array) SetParent(_ pdf.Object) {
	panic("not implemented") // TODO: Implement
}

func (arr *Array) Parent() pdf.Object {
	panic("not implemented") // TODO: Implement
}

func (arr *Array) Kind() pdf.ObjectKind {
	panic("not implemented") // TODO: Implement
}

func (arr *Array) Copy() (pdf.Object, error) {
	panic("not implemented") // TODO: Implement
}

func (arr *Array) Document() *pdf.Document {
	panic("not implemented") // TODO: Implement
}

func (arr *Array) GetIndirectReference() *pdf.Reference {
	panic("not implemented") // TODO: Implement
}

package pdf

type Dictionary struct{}

// MarshalPDF encodes the receiver a PDF bytes.
func (dict *Dictionary) MarshalPDF(_ *pdf.Writer) error {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) SetParent(_ pdf.Object) {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) Parent() pdf.Object {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) Kind() pdf.ObjectKind {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) Copy() (pdf.Object, error) {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) Document() *pdf.Document {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) GetIndirectReference() *pdf.Reference {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) Key(name Name) Object {
	panic("not implemented") // TODO: Implement
}

func (dict *Dictionary) AddKeyIndirect(name Name, obj Object) {
	panic("not implemented") // TODO: Implement
}

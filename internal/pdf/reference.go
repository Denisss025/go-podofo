package pdf

type Generation uint16

const FirstGeneration Generation = 1

type Reference struct {
	BaseObject

	objNo int
	genNo Generation
}

func NewReference(objecNo int, generationNo Generation) *Reference {
	return &Reference{objNo: objecNo, genNo: generationNo}
}

func (ref *Reference) Kind() ObjectKind {
	return ObjectKindReference
}

func (ref *Reference) Copy() (Object, error) {
	panic("not implemented") // TODO: implement me
}

func (ref *Reference) MarshalPDF(w *Writer) error {
	panic("not implemented") // TODO: implement me
}

func (ref *Reference) NumObjects() int {
	panic("not implemented") // TODO: implement me
}

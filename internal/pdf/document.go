package pdf

type Document struct {
	Objects []Object
}

func NewDocument() *Document { return new(Document) }

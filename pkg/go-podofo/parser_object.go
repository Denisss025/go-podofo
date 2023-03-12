package podofo

type ParserObject struct {
	Object
	Encrypt Encrypt
}

type ParserObjectOption func(*ParserObject)

func WithDocument(doc *Document) ParserObjectOption {
	panic("implement me") // TODO: implement me
}

func WithReference(ref *Reference) ParserObjectOption {
	panic("implement me") // TODO: implement me
}

func DelayedLoad(delayed bool) ParserObjectOption {
	panic("implement me") // TODO: implement me
}

func NewParserObject(r Reader, offset int64, options ...ParserObjectOption) (*ParserObject, error) {
	panic("not implemented") // TODO: implement me
}

func (obj *ParserObject) IsDictionary() bool {
	panic("not implemented") // TODO: implement me
}

func (obj *ParserObject) Dictionary() *Dictionary {
	panic("not implemented") // TODO: implement me
}

func (obj *ParserObject) Parse() error {
	panic("not implemented") // TODO: implement me
}

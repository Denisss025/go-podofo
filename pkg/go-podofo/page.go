package podofo

type Page struct{}

func (page *Page) Document() *Document {
	panic("not implemented") // TODO: implement me
}

func (page *Page) Object() Object {
	panic("not implemented") // TODO: implement me
}

func (page *Page) Rect() Rect {
	panic("not implemented") // TODO: implement me
}

type Pages struct{}

func (pg *Pages) Count() int {
	panic("not implemented") // TODO: implement me
}

func (pg *Pages) Add(size PageSize) *Page {
	panic("not implemented") // TODO: implement me
}

func (pg *Pages) Index(index int) *Page {
	panic("not implemented") // TODO: implement me
}

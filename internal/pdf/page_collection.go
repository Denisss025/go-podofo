package pdf

// TODO: implement

type PageCollection struct {
	element *Element
}

func NewPageCollection(document *Document) *PageCollection {
	panic("not implemented") // TODO: implement me
}

func (pc *PageCollection) Len() int {
	return getChildCount(pc.element.Object())
}

func getChildCount(node Object) int {
	countObj := node.Dictionary().Key(KeyCount)
	if countObj == nil {
		return 0
	}

	return countObj.(*Number).Int()
}

func (pc *PageCollection) Index(i int) (*Page, bool) {
	panic("not implemented") // TODO: implement me
}

func (pc *PageCollection) AddPage(size Rect) *Page {
	panic("not implemented") // TODO: implement me
}

func (pc *PageCollection) AddPageAt(index int, size Rect) (*Page, error) {
	panic("not implemented") // TODO: implement me
}

func (pc *PageCollection) AppendDocumentPages(
	document *Document, pageIndex int, count int,
) error {
	panic("not implemented") // TODO: implement me
}

func (pc *PageCollection) InsertDocumentPages(
	index int, document *Document, pageIndex int, count int,
) error {
	panic("not implemented") // TODO: implement me
}

func (pc *PageCollection) RemovePage(index int) bool {
	panic("not implemented") // TODO: implement me
}

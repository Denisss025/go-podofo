package pdf

type Element struct {
	object Object
}

func (e *Element) Object() Object { return e.object }

func (e *Element) Document() *Document { return e.object.Document() }

type DictionaryElement struct{ Element }

func (e *DictionaryElement) Dictionary() *Dictionary {
	return e.object.(*Dictionary)
}

type ArrayElement struct{ Element }

func (e *ArrayElement) Array() *Array { return e.object.(*Array) }

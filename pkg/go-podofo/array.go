package podofo

type Array struct {
	Objects []Object
}

func (array *Array) Kind() ObjectKind { return ObjectKindArray }

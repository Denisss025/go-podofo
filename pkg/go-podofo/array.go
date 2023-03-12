package podofo

type Array struct{}

func (array *Array) Kind() ObjectKind { return ObjectKindArray }

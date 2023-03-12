package podofo

type String string

func (s String) Kind() ObjectKind { return ObjectKindString }

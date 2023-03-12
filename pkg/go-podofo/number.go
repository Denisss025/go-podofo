package podofo

type Number string

func (num Number) Kind() ObjectKind { return ObjectKindNumber }

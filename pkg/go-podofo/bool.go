package podofo

type Bool bool

func (b Bool) Kind() ObjectKind { return ObjectKindBool }

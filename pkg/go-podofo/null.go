package podofo

type Null struct{}

func (null Null) Kind() ObjectKind { return ObjectKindNull }

package podofo

type Dictionary struct{}

func (d *Dictionary) Kind() ObjectKind {
	return ObjectKindDictionary
}

func (d *Dictionary) Key(name Name) Object {
	panic("not implemented") // TODO: implement me
}

func FindKey[T any](d *Dictionary, name Name, defval T) T {
	panic("not implemented") // TODO: implement me
}

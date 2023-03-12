package podofo

type RawData struct{}

func (data *RawData) Kind() ObjectKind { return ObjectKindRawData }

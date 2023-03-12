package podofo

type XRefEntryType int8

const (
	XRefEntryTypeUnknown XRefEntryType = iota - 1
	XRefEntryTypeFree
	XRefEntryTypeInUse
	XRefEntryTypeCompressed
)

func (t XRefEntryType) Byte() byte {
	switch t {
	case XRefEntryTypeFree:
		return 'f'
	case XRefEntryTypeInUse:
		return 'n'
	}

	// Unknown, Compressed, ...
	return 0
}

type XRefEntry struct {
	Offset     int64
	ObjectNum  int
	Type       XRefEntryType
	Parsed     bool
	Generation Generation
}

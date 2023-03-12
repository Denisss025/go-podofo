package podofo

type XRefEntry struct {
	Entry  xrefEntry
	Parsed bool
}

type xrefEntry interface{ xrefEntry() }

type XRefEntryFree struct {
	ObjectNumber int
	Generation   Generation
}

func (f XRefEntryFree) xrefEntry() {}

type XRefEntryInUse struct {
	Offset     int64
	Generation Generation
}

func (n XRefEntryInUse) xrefEntry() {}

type XRefEntryCompressed struct {
	ObjectNumber int
	Index        int64
}

func (c XRefEntryCompressed) xrefEntry() {}

package podofo

// TODO: implement DataContainer

type DataContainer interface {
	Owner() Object
	SetOwner(owner Object)

	ObjectDocument() *Document
	GetIndirectObject(*Reference) Object

	SetDirty()
	ResetDirty()

	IsIndirectReferenceAllowed(obj Object) bool
}

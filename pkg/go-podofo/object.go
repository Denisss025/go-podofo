package podofo

import "io"

type ObjectKind uint8

const (
	// ObjectKindUnknown
	ObjectKindUnknown ObjectKind = iota
	// ObjectKindBool
	ObjectKindBool
	// ObjectKindNumber
	ObjectKindNumber
	// ObjectKindString
	ObjectKindString
	// ObjectKindName
	ObjectKindName
	// ObjectKindArray
	ObjectKindArray
	// ObjectKindDictionary
	ObjectKindDictionary
	// ObjectKindNull
	ObjectKindNull
	// ObjectKindReference
	ObjectKindReference
	// ObjectKindRawData
	ObjectKindRawData
)

type Object interface {
	io.WriterTo

	Kind() ObjectKind
}

// func (obj *Object) Kind() ObjectKind {
// 	panic("not implemented") // TODO: implement me
// }

// func (obj *Object) Reference() *Reference {
// 	panic("not implemented") // TODO: implement me
// }

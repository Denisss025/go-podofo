package pdf

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
	Marshaler

	SetParent(Object)
	Parent() Object

	Kind() ObjectKind
	Copy() (Object, error)
	Document() *Document
	Dictionary() *Dictionary
	GetIndirectReference() *Reference
}

type BaseObject struct {
	parent   Object
	variant  Variant
	document *Document
	indirect *Reference

	reader io.ReadSeeker

	isDelayedLoadDone       bool
	isDelayedLoadStreamDone bool
}

func (obj *BaseObject) Copy() (BaseObject, error) {
	retval := BaseObject{
		variant:                 obj.variant,
		isDelayedLoadDone:       true,
		isDelayedLoadStreamDone: true,
		reader:                  obj.reader,
	}

	err := obj.DelayedLoadStream()
	obj.reader = nil

	retval.setVariantOwner()

	return retval, err
}

func (obj *BaseObject) setVariantOwner() {
	panic("not implemented") // TODO: implement me
}

func (obj *BaseObject) DelayedLoadStream() error {
	panic("not implemented") // TODO: implement me
}

func (obj *BaseObject) SetParent(parent Object) {
	obj.parent = parent
}

func (obj *BaseObject) Parent() Object { return obj.parent }

func (obj *BaseObject) Document() *Document {
	return obj.document
}

func (obj *BaseObject) GetIndirectReference() *Reference {
	return obj.indirect
}

// func (obj *Object) Kind() ObjectKind {
// 	panic("not implemented") // TODO: implement me
// }

// func (obj *Object) Reference() *Reference {
// 	panic("not implemented") // TODO: implement me
// }

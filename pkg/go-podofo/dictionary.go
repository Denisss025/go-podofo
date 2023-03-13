package podofo

import "io"

// Dictionary is the PDF Dictionary data type.
type Dictionary struct {
	keys map[Name]Object
}

// Kind returns the kind of the PDFObject.
func (d *Dictionary) Kind() ObjectKind {
	return ObjectKindDictionary
}

func (d *Dictionary) WriteTo(w io.Writer) (n int64, err error) {
	panic("not implemented") // TODO: implement me
}

// AddKey adds key to the dictionary.
func (d *Dictionary) AddKey(name Name, obj Object) {
	d.keys[name] = obj
}

// AddKeyIndirect adds key to the dictionary.
func (d *Dictionary) AddKeyIndirect(name Name, obj Object) error {
	panic("not implemented") // TODO: implement me
}

// Key finds an object in the dictionry. Nil is returned if the
// key does not exist.
func (d *Dictionary) Key(name Name) Object {
	return d.keys[name]
}

// Int finds an object and converts it to int64.
func (d *Dictionary) Int(name Name, defval int64) int64 {
	panic("not implemented") // TODO: implement me
}

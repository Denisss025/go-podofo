package podofo

import (
	"github.com/denisss025/go-podofo/internal/pdf"
)

// Dictionary is the PDF Dictionary data type.
type Dictionary struct {
	pdf.BaseObject
	// TODO? pdf.DataContainer?

	keys map[Name]Object
}

// Kind returns the kind of the PDFObject.
func (d *Dictionary) Kind() ObjectKind {
	return pdf.ObjectKindDictionary
}

func (d *Dictionary) MarshalPDF(w *pdf.Writer) error {
	panic("not implemented") // TODO: implement me
}

// AddKey adds key to the dictionary.
func (d *Dictionary) AddKey(name Name, obj Object) {
	// TODO? need copy?
	d.keys[name] = obj
}

func (d *Dictionary) Copy() (Object, error) {
	panic("not implemented") // TODO: implement me
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

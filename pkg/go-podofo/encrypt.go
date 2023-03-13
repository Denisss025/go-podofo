package podofo

type Encrypt struct{}

func EncryptFromObject(obj Object) (*Encrypt, error) {
	panic("not implemented") // TODO: implement me
}

func (enc *Encrypt) Authenticate(password string, objectID *String) bool {
	panic("not implemented") // TODO: implement me
}

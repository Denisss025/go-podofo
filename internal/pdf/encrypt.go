package pdf

// TODO: implement Encryptor

type Encrypt interface{}

type encrypt struct{}

func NewEncrypt() Encrypt { return new(encrypt) }

package podofo

type Tokenizer struct {
	buffer []byte
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{buffer: make([]byte, BufferSize)}
}

func (tok *Tokenizer) TryReadNextToken(r Reader) (token []byte, err error) {
	panic("not implemented") // TODO: implement me
}

func (tok *Tokenizer) ReadNextNumber(r Reader) (number int64, err error) {
	panic("not implemented") // TODO: implement me
}

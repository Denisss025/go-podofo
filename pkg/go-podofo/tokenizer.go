package podofo

type Tokenizer struct {
	buffer []byte
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{buffer: make([]byte, BufferSize)}
}

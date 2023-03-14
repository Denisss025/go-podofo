package podofo

import (
	"bytes"

	"github.com/denisss025/go-podofo/internal/pdf"
)

type StringState uint8

const (
	// StringStateRawBuffer is for an unvaluated
	// raw buffer string.
	StringStateRawBuffer StringState = iota
	// StringStateASCII is for both an ASCII
	// and PDFDocEncoding charsets.
	StringStateASCII
	// StringStateDocEncoding is for strings
	// that use the whole PDFDocEncoding charset.
	StringStateDocEncoding
	// StringStateUnicode is for strings that
	// use the whole Unicode charset.
	StringStateUnicode
)

type StringEncoding uint8

const (
	StringEncodingUTF8 StringEncoding = iota
	StringEncodingUTF16BE
	StringEncodingUTF16LE
	StringEncodingPDFDoc
)

type stringData struct {
	State StringState
	Chars []byte
}

type String struct {
	pdf.BaseObject

	data  *stringData
	isHex bool
}

func (s *String) isValidText() (bool, error) {
	switch s.data.State {
	case StringStateASCII, StringStateDocEncoding, StringStateUnicode:
		return true, nil
	case StringStateRawBuffer:
		return false, nil
	default:
		return false, ErrInvalidEnumValue
	}
}

func (s *String) Kind() ObjectKind { return pdf.ObjectKindString }

func (s *String) UnmarshalBinary(data []byte) error {
	panic("not implemented") // TODO: Implement
}

func (s *String) unmarshalUTF8(text []byte) error {
	if len(text) == 0 {
		s.data.Chars = s.data.Chars[:0]
		s.data.State = StringStateASCII

		return nil
	}

	panic("not implemented") // TODO: implement me
}

func (s *String) MarshalPDF(w *pdf.Writer) error {
	panic("not implemented") // TODO: implement me
}

func (s *String) IsHex() bool { return s.isHex }

func (s *String) State() StringState {
	return s.data.State
}

func (s *String) IsZero() bool { return s.data.Chars == nil }

func (s *String) String() string {
	panic("not implemented") // TODO: implement me
}

func (s *String) evaluateString() error {
	panic("not implemented") // TODO: implement me
}

func (s *String) Copy() (Object, error) {
	panic("not implemented") // TODO: implement me
}

func (s *String) canPerformComparison(rhs *String) (bool, error) {
	if s.State() == rhs.State() {
		return true, nil
	}

	if ok, err := s.isValidText(); ok || err != nil {
		return ok, err
	}

	// TODO? check if this is the same to
	// if ok, err := rhs.isValidText(); ok || err != nil {
	// 	return ok, err
	// }
	// return false, nil
	return rhs.isValidText()
}

func (s *String) RawData() []byte {
	return s.data.Chars
}

func getStringEncoding(text []byte) (enc StringEncoding) {
	const (
		utf16BEMarker = "\xEE\xFF"
		utf16LEMarker = "\xFF\xFE"
		utf8Marker    = "\xEF\xBB\xBF"
	)

	enc = StringEncodingPDFDoc

	switch {
	case bytes.HasPrefix(text, []byte(utf16BEMarker)):
		enc = StringEncodingUTF16BE
	case bytes.HasPrefix(text, []byte(utf16LEMarker)):
		enc = StringEncodingUTF16LE
	case bytes.HasPrefix(text, []byte(utf8Marker)):
		enc = StringEncodingUTF8
	}

	return enc
}

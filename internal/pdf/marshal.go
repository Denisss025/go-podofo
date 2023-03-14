package pdf

import "io"

// WriteFlag is the flag that is used by the pdf.Writer.
type WriteFlag uint16

const (
	// WriteFlagClean is used to created a PDF
	// that is readable in a text editor, i.e.
	// isert spaces and linebreaks between tokens.
	WriteFlagClean WriteFlag = 1 << iota
	// WriteFlagNoInlineLiteral is used to prevent writing spaces before literal types
	// (numerical, references, null).
	WriteFlagNoInlineLiteral
	// WriteFlagNoFlateCompress is used to write
	// PDF with Flate compression.
	WriteFlagNoFlateCompress
	// WriteFlagNoPDFAPreserve is used to write
	// compact (WriteFlagsClean is unsed) code,
	// preserving PDF/A compliance is not required.
	WriteFlagNoPDFAPreserve WriteFlag = 256
)

type WriterOptionFunc func(*Writer)

func WriteClean() WriterOptionFunc {
	return func(w *Writer) { w.flags |= WriteFlagClean }
}

func WriteNoInlineLiteral() WriterOptionFunc {
	return func(w *Writer) { w.flags |= WriteFlagNoInlineLiteral }
}

func WriteNoCompress() WriterOptionFunc {
	return func(w *Writer) { w.flags |= WriteFlagNoFlateCompress }
}

func WriteNoPDFAPreserve() WriterOptionFunc {
	return func(w *Writer) { w.flags |= WriteFlagNoPDFAPreserve }
}

func WriteEncryptor(encrypt *Encrypt) WriterOptionFunc {
	return func(w *Writer) { w.encrypt = encrypt }
}

type Writer struct {
	out     io.Writer
	encrypt Encrypt
	flags   WriteFlag
}

func NewWriter(w io.Writer, options ...WriterOptionFunc) *Writer {
	wx := &Writer{
		out:     w,
		encrypt: NewEncrypt(), // nop-encrypt by default
	}

	for _, opt := range options {
		opt(wx)
	}

	return wx
}

func (w *Writer) IsCleanWrite() bool { return w.flags&WriteFlagClean != 0 }

func (w *Writer) Write(p []byte) (n int, err error) {
	panic("not implemented") // TODO: implement me
}

// Marshaler is the interface implemented by PDF objects that can marshal
// themselves into valid PDF stream.
type Marshaler interface {
	// MarshalPDF encodes the receiver a PDF bytes.
	MarshalPDF(*Writer) error
}

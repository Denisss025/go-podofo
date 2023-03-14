package podofo

import (
	"io"

	"github.com/denisss025/go-podofo/internal/pdf"
)

func write(w io.Writer, p []byte) error {
	_, err := w.Write(p)

	return err
}

func writeClean(w *pdf.Writer, p []byte) (err error) {
	if w.IsCleanWrite() {
		_, err = w.Write(p)
	}

	return err
}

func writeString(w io.Writer, s string) error {
	return write(w, []byte(s))
}

func writeStringClean(w *pdf.Writer, s string) error {
	return writeClean(w, []byte(s))
}

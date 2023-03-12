package podofo

import "fmt"

// PDFVersion is an enumeration for supported PDF versions.
type PDFVersion string

const (
	// PDFVersion10 PDF 1.0.
	PDFVersion10 PDFVersion = "1.0"
	// PDFVersion11 PDF 1.1.
	PDFVersion11 PDFVersion = "1.1"
	// PDFVersion12 PDF 1.2.
	PDFVersion12 PDFVersion = "1.2"
	// PDFVersion13 PDF 1.3.
	PDFVersion13 PDFVersion = "1.3"
	// PDFVersion14 PDF 1.4.
	PDFVersion14 PDFVersion = "1.4"
	// PDFVersion15 PDF 1.5.
	PDFVersion15 PDFVersion = "1.5"
	// PDFVersion16 PDF 1.6.
	PDFVersion16 PDFVersion = "1.6"
	// PDFVersion17 PDF 1.7.
	PDFVersion17 PDFVersion = "1.7"
	// PDFVersion20 PDF 2.0
	PDFVersion20 PDFVersion = "2.0"
)

func (ver PDFVersion) Validate() error {
	switch ver {
	case PDFVersion10,
		PDFVersion11,
		PDFVersion12,
		PDFVersion13,
		PDFVersion14,
		PDFVersion15,
		PDFVersion16,
		PDFVersion17,
		PDFVersion20:
		return nil
	}

	return fmt.Errorf("%w: %s", ErrUnsupportedVersion, string(ver))
}

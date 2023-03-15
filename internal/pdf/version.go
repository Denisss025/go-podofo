package pdf

import (
	"errors"
	"fmt"
)

// Version is an enumeration for supported PDF versions.
type Version string

var ErrUnsupportedVersion = errors.New("unsupported pdf version")

const (
	// Version10 PDF 1.0.
	Version10 Version = "1.0"
	// Version11 PDF 1.1.
	Version11 Version = "1.1"
	// Version12 PDF 1.2.
	Version12 Version = "1.2"
	// Version13 PDF 1.3.
	Version13 Version = "1.3"
	// Version14 PDF 1.4.
	Version14 Version = "1.4"
	// Version15 PDF 1.5.
	Version15 Version = "1.5"
	// Version16 PDF 1.6.
	Version16 Version = "1.6"
	// Version17 PDF 1.7.
	Version17 Version = "1.7"
	// Version20 PDF 2.0
	Version20 Version = "2.0"
)

func (ver Version) Validate() error {
	switch ver {
	case Version10,
		Version11,
		Version12,
		Version13,
		Version14,
		Version15,
		Version16,
		Version17,
		Version20:
		return nil
	}

	return fmt.Errorf("%w: %s", ErrUnsupportedVersion, string(ver))
}

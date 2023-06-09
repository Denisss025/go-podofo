package podofo

import (
	"errors"
	"io"
	"io/fs"

	"github.com/denisss025/go-podofo/internal/pdf"
)

var (
	ErrUnknown                   = errors.New("unknown error")
	ErrInvalidHandle             = errors.New("invalid handle")
	ErrFileNotFound              = fs.ErrNotExist
	ErrInvalidDeviceOperation    = errors.New("invalid device operation")
	ErrUnexpectedEOF             = io.ErrUnexpectedEOF
	ErrValueOutOfRange           = pdf.ErrValueOutOfRange
	ErrInternalLogic             = errors.New("internal logic")
	ErrInvalidEnumValue          = errors.New("invalid enum value")
	ErrBrokenFile                = errors.New("file is broken")
	ErrPageNotFound              = errors.New("page not found")
	ErrNoPDFFile                 = errors.New("not a PDF file")
	ErrNoXRef                    = errors.New("no valid XRef")
	ErrNoTrailer                 = errors.New("no trailer")
	ErrNoNumber                  = errors.New("not a number")
	ErrNoObject                  = errors.New("not an object")
	ErrNoEOFToken                = errors.New("EOF token not found")
	ErrInvalidTrailerSize        = errors.New("invalid trailer size")
	ErrInvalidDataType           = errors.New("invalid data type")
	ErrInvalidXRef               = errors.New("invalid XRef")
	ErrInvalidXRefStream         = errors.New("invalid XRef stream")
	ErrInvalidXRefType           = errors.New("invalid XRef type")
	ErrInvalidPredictor          = errors.New("invalid predictor")
	ErrInvalidStrokeStyle        = errors.New("invalid stroke style")
	ErrInvalidHexString          = errors.New("invalid hex string")
	ErrInvalidStream             = errors.New("invalid stream")
	ErrInvalidStreamLength       = errors.New("invalid stream len")
	ErrInvalidKey                = errors.New("invalid key")
	ErrInvalidName               = errors.New("invalid name")
	ErrInvalidEncryptionDict     = errors.New("invalid encryption dict")
	ErrInvalidPassword           = errors.New("invalid password")
	ErrInvalidFontData           = errors.New("invalid font data")
	ErrInvalidContentStream      = errors.New("invalid content stream")
	ErrUnsupportedVersion        = pdf.ErrUnsupportedVersion
	ErrUnsupportedFilter         = errors.New("unsupported filter")
	ErrUnsupportedFontFormat     = errors.New("unsupported font format")
	ErrUnsupportedImageFormat    = errors.New("unsupported image format")
	ErrActionAlreadyPresent      = errors.New("action already present")
	ErrWrongDestinationType      = errors.New("wrong destination type")
	ErrMissingEndStream          = errors.New("missing steram end")
	ErrDate                      = errors.New("bad date")
	ErrFlate                     = errors.New("flate")
	ErrFreeType                  = errors.New("free type")
	ErrSignature                 = errors.New("signature")
	ErrCannotConvertColor        = pdf.ErrCannotConvertColor
	ErrNotImplemented            = pdf.ErrNotImplemented
	ErrDestinationAlreadyPresent = errors.New("destination already present")
	ErrChangeOnImmutable         = errors.New("change on immutable")
	ErrNotCompiled               = errors.New("not compiled")
	ErrOutlineItemAlreadyPresent = errors.New("outline already present")
	ErrNotLoadedForUpdate        = errors.New("not loaded for update")
	ErrCannotEncrypUpdate        = errors.New("cannot encrypt update")
	ErrXMPMetadata               = errors.New("xmp metadata")
)

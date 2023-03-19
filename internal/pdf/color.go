package pdf

import (
	"errors"
	"fmt"
	"image/color"
	"strconv"
	"unicode"

	"golang.org/x/exp/constraints"
)

// ColorSpace is an enum for the colorspaces supported by PDF.
type ColorSpace uint8

const (
	// ColorSpaceUnknown is an uknown colorspace.
	ColorSpaceUnknown ColorSpace = iota
	ColorSpaceDeviceGray
	ColorSpaceDeviceRGB
	ColorSpaceDeviceCMYK
	ColorSpaceCalGray
	ColorSpaceCalRGB
	ColorSpaceLab ///< CIE-Lab
	ColorSpaceICCBased
	ColorSpaceIndexed
	ColorSpacePattern
	ColorSpaceSeparation
	ColorSpaceDeviceN
)

type Color interface {
	ColorSpace() ColorSpace
	Validate() error
}

type GrayColor color.Gray16

func (gray GrayColor) ColorSpace() ColorSpace { return ColorSpaceDeviceGray }

func (gray GrayColor) Validate() error { return nil }

func (gray GrayColor) RGBA() (r, g, b, a uint32) {
	return color.Gray16(gray).RGBA()
}

func (gray *GrayColor) UnmarshalText(text []byte) error {
	panic("not implemented") // TODO: Implement
}

type RGBColor color.NRGBA64

func (rgb RGBColor) ColorSpace() ColorSpace { return ColorSpaceDeviceRGB }

func (rgb RGBColor) Validate() error { return nil }

func (rgb RGBColor) RGBA() (r, g, b, a uint32) {
	return color.NRGBA64(rgb).RGBA()
}

func (rgb *RGBColor) UnmarshalText(text []byte) error {
	const (
		hexInt = 16
		rgbLen = 7
		sharp  = '#'

		rshift = 16
		gshift = 8
		bshift = 0

		mask      = 0xFF
		maxUInt16 = 0xFFFF
	)

	if len(text) != rgbLen || text[0] != sharp {
		return fmt.Errorf("unmsarshal RBB color: %w", ErrCannotConvertColor)
	}

	n, err := strconv.ParseUint(string(text[1:]), hexInt, strconv.IntSize)
	if err != nil {
		return fmt.Errorf("unmarshal RGB color: %w", err)
	}

	rgb.R = uint16(uint64(
		(n>>rshift)&mask) * uint64(maxUInt16) / uint64(mask))
	rgb.G = uint16(uint64(
		(n>>gshift)&mask) * uint64(maxUInt16) / uint64(mask))
	rgb.B = uint16(uint64(
		(n>>bshift)&mask) * uint64(maxUInt16) / uint64(mask))

	return nil
}

type CMYKColor color.CMYK

func (cmyk CMYKColor) ColorSpace() ColorSpace { return ColorSpaceDeviceCMYK }

func (cmyk CMYKColor) Validate() error { return nil }

func (cmyk CMYKColor) RGBA() (r, g, b, a uint32) {
	return color.CMYK(cmyk).RGBA()
}

func (cmyk *CMYKColor) UnmarshalText(text []byte) error {
	panic("not implemented") // TODO: Implement
}

type CieLabColor struct {
	// CieL indicates lightness.
	CieL float64
	// CieA indicates A color value.
	CieA float64
	// CieB  indicates B color value.
	CieB float64
}

func (lab CieLabColor) ColorSpace() ColorSpace { return ColorSpaceLab }

func (lab CieLabColor) Validate() error {
	const (
		lMax = 100.0

		colorMin = -128.0
		colorMax = 127.0
	)

	return errors.Join(
		checkRange(lab.CieL, 0.0, lMax),
		checkRange(lab.CieA, colorMin, colorMax),
		checkRange(lab.CieB, colorMin, colorMax),
	)
}

type SeparationColor struct {
	Name           string
	Density        float64
	AlternateColor Color
}

func (sep SeparationColor) ColorSpace() ColorSpace { return ColorSpaceSeparation }

func (sep SeparationColor) Validate() error {
	return sep.AlternateColor.Validate()
}

func ParseColor(text []byte) (color Color, err error) {
	const (
		dot        = '.'
		sharp      = '#'
		arrayStart = '['

		rgbLen  = 7
		cmykLen = 9
	)

	switch {
	case len(text) == 0:
		color = GrayColor{}
	case text[0] == dot || unicode.IsDigit(rune(text[0])):
		gray := GrayColor{}
		err = gray.UnmarshalText(text)
		color = gray
	case text[0] == sharp && len(text) == rgbLen:
		rgb := RGBColor{}
		err = rgb.UnmarshalText(text)
		color = rgb
	case text[0] == sharp && len(text) == cmykLen:
		cmyk := CMYKColor{}
		err = cmyk.UnmarshalText(text)
		color = cmyk
	case text[0] == arrayStart:
		// TODO: array
		panic("not implemented") // TODO: implement me
	default:
		// TODO: named color
		panic("not implemented") // TODO: implement me
	}

	return color, err
}

func SeparationColorNone() SeparationColor {
	return SeparationColor{
		Name:           "None",
		AlternateColor: CMYKColor{},
	}
}

func SeparationColorAll() SeparationColor {
	const (
		maxUInt8 uint8   = 0xFF
		density  float64 = 1.0
	)

	return SeparationColor{
		Name:    "All",
		Density: density,
		AlternateColor: CMYKColor{
			C: maxUInt8,
			M: maxUInt8,
			Y: maxUInt8,
			K: maxUInt8,
		},
	}
}

func TransparentColor() Color {
	panic("not implemented") // TODO: implement me
}

func ColorFromObject(obj Object) (Color, error) {
	panic("not implemented") // TODO: implement me
}

func ConvertToGrayScale(inColor Color) (gray GrayColor, err error) {
	switch c := inColor.(type) {
	case GrayColor:
		return c, nil
	case RGBColor:
		gray = GrayColor(color.Gray16Model.Convert(c).(color.Gray16))
	case CMYKColor:
		rgb, err := ConvertToRGB(c)
		if err != nil {
			return gray, err
		}

		return ConvertToGrayScale(rgb)
	case SeparationColor:
		cmyk, ok := c.AlternateColor.(CMYKColor)
		if !ok {
			return gray, ErrNotImplemented
		}

		return ConvertToGrayScale(cmyk)
	default:
		err = ErrCannotConvertColor
	}

	return gray, err
}

func ConvertToRGB(inColor Color) (rgb RGBColor, err error) {
	switch c := inColor.(type) {
	case RGBColor:
		return c, nil
	case GrayColor:
		rgb = RGBColor(color.NRGBA64Model.Convert(c).(color.NRGBA64))
	case CMYKColor:
		rgb = RGBColor(color.NRGBA64Model.Convert(c).(color.NRGBA64))
	case SeparationColor:
		cmyk, ok := c.AlternateColor.(CMYKColor)
		if !ok {
			return rgb, ErrNotImplemented
		}

		return ConvertToRGB(cmyk)
	default:
		err = ErrCannotConvertColor
	}

	return rgb, err
}

func ConvertToCMYK(inColor Color) (cmyk CMYKColor, err error) {
	switch c := inColor.(type) {
	case CMYKColor:
		return c, nil
	case GrayColor:
		cmyk = CMYKColor(color.CMYKModel.Convert(c).(color.CMYK))
	case RGBColor:
		cmyk = CMYKColor(color.CMYKModel.Convert(c).(color.CMYK))
	default:
		err = ErrCannotConvertColor
	}

	return cmyk, err
}

func checkRange[N constraints.Ordered](val, min, max N) error {
	if val < min || val > max {
		return ErrValueOutOfRange
	}

	return nil
}

// TODO: list of PdfColor methods:
// - ToArray() const -> PdfArray;
// - BuildColorSpace(PdfDocument&) -> PdfObject*.

// TODO: list of PdfColor private fields:
// - bool isTransparent;
// - PdfColorSpace colorSpace;

// TODO: static private field:
// - const unsigned* const hexDigitMap.

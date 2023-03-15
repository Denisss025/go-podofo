package pdf

import (
	"errors"
	"fmt"
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

type GrayColor struct {
	Value float64
}

func (gray GrayColor) ColorSpace() ColorSpace { return ColorSpaceDeviceGray }

func (gray GrayColor) Validate() error {
	const colorMax = 1.0

	return checkRange(gray.Value, 0.0, colorMax)
}

func (gray *GrayColor) UnmarshalText(text []byte) error {
	panic("not implemented") // TODO: Implement
}

type RGBColor struct {
	Red   float64
	Green float64
	Blue  float64
}

func (rgb RGBColor) ColorSpace() ColorSpace { return ColorSpaceDeviceRGB }

func (rgb RGBColor) Validate() error {
	const colorMax = 1.0

	return errors.Join(
		checkRange(rgb.Red, 0.0, colorMax),
		checkRange(rgb.Green, 0.0, colorMax),
		checkRange(rgb.Blue, 0.0, colorMax),
	)
}

func (rgb *RGBColor) UnmarshalText(text []byte) error {
	const (
		hexInt = 16
		rgbLen = 7
		sharp  = '#'

		rshift = 16
		gshift = 8
		bshift = 0

		mask = 0xFF
	)

	if len(text) != rgbLen || text[0] != sharp {
		return fmt.Errorf("unmsarshal RBB color: %w", ErrCannotConvertColor)
	}

	n, err := strconv.ParseUint(string(text[1:]), hexInt, strconv.IntSize)
	if err != nil {
		return fmt.Errorf("unmarshal RGB color: %w", err)
	}

	rgb.Red = float64((n>>rshift)&mask) / float64(mask)
	rgb.Green = float64((n>>gshift)&mask) / float64(mask)
	rgb.Blue = float64((n>>bshift)&mask) / float64(mask)

	return nil
}

type CMYKColor struct {
	Cyan    float64
	Magenta float64
	Yellow  float64
	Black   float64
}

func (cmyk CMYKColor) ColorSpace() ColorSpace { return ColorSpaceDeviceCMYK }

func (cmyk CMYKColor) Validate() error {
	const colorMax = 1.0

	return errors.Join(
		checkRange(cmyk.Cyan, 0.0, colorMax),
		checkRange(cmyk.Magenta, 0.0, colorMax),
		checkRange(cmyk.Yellow, 0.0, colorMax),
		checkRange(cmyk.Black, 0.0, colorMax),
	)
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
	const color = 1.0

	return SeparationColor{
		Name:    "All",
		Density: color,
		AlternateColor: CMYKColor{
			Cyan:    color,
			Magenta: color,
			Yellow:  color,
			Black:   color,
		},
	}
}

func TransparentColor() Color {
	panic("not implemented") // TODO: implement me
}

func ColorFromObject(obj Object) (Color, error) {
	panic("not implemented") // TODO: implement me
}

func ConvertToGrayScale(color Color) (gray GrayColor, err error) {
	const (
		rweight = 0.299
		gweight = 0.587
		bweignt = 0.114
	)

	switch color := color.(type) {
	case GrayColor:
		return color, nil
	case RGBColor:
		gray.Value = rweight*color.Red +
			bweignt*color.Blue +
			gweight*color.Green
	case CMYKColor:
		rgb, err := ConvertToRGB(color)
		if err != nil {
			return gray, err
		}

		return ConvertToGrayScale(rgb)
	case SeparationColor:
		cmyk, ok := color.AlternateColor.(CMYKColor)
		if !ok {
			return gray, ErrNotImplemented
		}

		return ConvertToGrayScale(cmyk)
	default:
		err = ErrCannotConvertColor
	}

	return gray, err
}

func ConvertToRGB(color Color) (rgb RGBColor, err error) {
	const colorMax = 1.0

	switch color := color.(type) {
	case RGBColor:
		return color, nil
	case GrayColor:
		rgb.Red = color.Value
		rgb.Blue = color.Value
		rgb.Green = color.Value
	case CMYKColor:
		rgb.Red = (colorMax - color.Cyan) * (colorMax - color.Black)
		rgb.Green = (colorMax - color.Magenta) * (colorMax - color.Black)
		rgb.Blue = (colorMax - color.Yellow) * (colorMax - color.Black)
	case SeparationColor:
		cmyk, ok := color.AlternateColor.(CMYKColor)
		if !ok {
			return rgb, ErrNotImplemented
		}

		return ConvertToRGB(cmyk)
	default:
		err = ErrCannotConvertColor
	}

	return rgb, err
}

func ConvertToCMYK(color Color) (cmyk CMYKColor, err error) {
	const colorMax = 1.0

	switch color := color.(type) {
	case CMYKColor:
		return color, nil
	case GrayColor:
		cmyk.Cyan = 0.0
		cmyk.Magenta = 0.0
		cmyk.Yellow = 0.0
		cmyk.Black = colorMax - color.Value
	case RGBColor:
		cmyk.Black = min(
			colorMax-color.Red,
			colorMax-color.Blue,
			colorMax-color.Green,
		)

		if cmyk.Black < colorMax {
			x := colorMax - cmyk.Black

			cmyk.Cyan = (x - color.Red) / x
			cmyk.Magenta = (x - color.Green) / x
			cmyk.Yellow = (x - color.Blue) / x
		}
	default:
		err = ErrCannotConvertColor
	}

	return cmyk, err
}

func min[N constraints.Ordered](a N, bcd ...N) N {
	for _, b := range bcd {
		if b < a {
			a = b
		}
	}

	return a
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

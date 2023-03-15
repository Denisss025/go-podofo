package pdf_test

import (
	"testing"

	"github.com/denisss025/go-podofo/internal/pdf"
	"github.com/stretchr/testify/assert"
)

func TestRGBColorUnmarshalText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		text    string
		expect  pdf.RGBColor
		wantErr bool
	}{
		{
			name:    "black",
			text:    "#000000",
			expect:  pdf.RGBColor{},
			wantErr: false,
		}, {
			name: "white",
			text: "#FFFFFF",
			expect: pdf.RGBColor{
				Red:   1.0,
				Green: 1.0,
				Blue:  1.0,
			},
			wantErr: false,
		}, {
			name: "#fcba03",
			text: "#fcba03",
			expect: pdf.RGBColor{
				Red:   252.0 / 255.0,
				Green: 186.0 / 255.0,
				Blue:  3.0 / 255.0,
			},
			wantErr: false,
		}, {
			name:    "cmyk",
			text:    "#00000000",
			expect:  pdf.RGBColor{},
			wantErr: true,
		}, {
			name:    "ZZZ",
			text:    "#ZZZZZZ",
			expect:  pdf.RGBColor{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rgb pdf.RGBColor

			if tt.wantErr {
				assert.Error(t, rgb.UnmarshalText([]byte(tt.text)))
			} else {
				assert.NoError(t, rgb.UnmarshalText([]byte(tt.text)))

				assert.Equal(t, tt.expect, rgb)
			}
		})
	}
}

func TestConvertToCMYK(t *testing.T) {
	t.Parallel()

	mapCmyk := func(cmyk pdf.CMYKColor) map[string]float64 {
		return map[string]float64{
			"cyan":    cmyk.Cyan,
			"magenta": cmyk.Magenta,
			"yellow":  cmyk.Yellow,
			"black":   cmyk.Black,
		}
	}

	tests := []struct {
		name     string
		color    pdf.Color
		wantCmyk pdf.CMYKColor
		wantErr  bool
	}{
		{
			name: "rgb2cmyk-#0ac77d",
			color: pdf.RGBColor{
				Red:   11.0 / 255.0,
				Green: 200.0 / 255.0,
				Blue:  127.0 / 255.0,
			},
			wantCmyk: pdf.CMYKColor{
				Cyan:    0.95,
				Magenta: 0.0,
				Yellow:  0.37,
				Black:   0.22,
			},
			wantErr: false,
		}, {
			name: "rgb2cmyk-#2b00ff",
			color: pdf.RGBColor{
				Red:   43.0 / 255.0,
				Green: 0.0,
				Blue:  1.0,
			},
			wantCmyk: pdf.CMYKColor{
				Cyan:    0.83,
				Magenta: 1.0,
				Yellow:  0.0,
				Black:   0.0,
			},
			wantErr: false,
		}, {
			name: "cmyk2cmyk",
			color: pdf.CMYKColor{
				Cyan:    0.1,
				Magenta: 0.2,
				Yellow:  0.3,
				Black:   0.4,
			},
			wantCmyk: pdf.CMYKColor{
				Cyan:    0.1,
				Magenta: 0.2,
				Yellow:  0.3,
				Black:   0.4,
			},
			wantErr: false,
		}, {
			name: "gray2cmyk",
			color: pdf.GrayColor{
				Value: 55.0 / 255.0,
			},
			wantCmyk: pdf.CMYKColor{
				Black: 0.78,
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotCmyk, err := pdf.ConvertToCMYK(tt.color)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.InDeltaMapValues(t, mapCmyk(tt.wantCmyk), mapCmyk(gotCmyk), 0.01)
				// assert.Equal(t, tt.wantCmyk, gotCmyk)
			}
		})
	}

	t.Run("gray to RGB to CMYK", func(t *testing.T) {
		t.Parallel()

		gray := pdf.GrayColor{Value: float64(0xCC) / 255.00}
		expectRGB := pdf.RGBColor{
			Red:   gray.Value,
			Green: gray.Value,
			Blue:  gray.Value,
		}
		expectCMYK := pdf.CMYKColor{
			Black: 0.2,
		}

		rgb, err := pdf.ConvertToRGB(gray)
		assert.NoError(t, err)
		assert.Equal(t, expectRGB, rgb)

		cmyk, err := pdf.ConvertToCMYK(rgb)
		assert.NoError(t, err)
		assert.InDeltaMapValues(t, mapCmyk(expectCMYK), mapCmyk(cmyk), 0.01)
	})
}

package pdf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/denisss025/go-podofo/internal/pdf"
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
				R: 0xFFFF,
				G: 0xFFFF,
				B: 0xFFFF,
			},
			wantErr: false,
		}, {
			name: "#fcba03",
			text: "#fcba03",
			expect: pdf.RGBColor{
				R: uint16(uint64(252 * 0xFFFF / 255)),
				G: uint16(uint64(186 * 0xFFFF / 255)),
				B: uint16(uint64(3 * 0xFFFF / 255)),
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
		const maxUInt8 = float64(uint8(0xFF))

		return map[string]float64{
			"cyan":    float64(cmyk.C) / maxUInt8,
			"magenta": float64(cmyk.M) / maxUInt8,
			"yellow":  float64(cmyk.Y) / maxUInt8,
			"black":   float64(cmyk.K) / maxUInt8,
		}
	}

	toCmyk := func(v float64) uint8 {
		const maxUInt8 = float64(uint8(0xFF))

		return uint8(v * maxUInt8)
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
				R: uint16(uint64(11 * 0xFFFF / 255)),
				G: uint16(uint64(200 * 0xFFFF / 255)),
				B: uint16(uint64(127 * 0xFFFF / 255)),
			},
			wantCmyk: pdf.CMYKColor{
				C: toCmyk(0.95),
				M: toCmyk(0.0),
				Y: toCmyk(0.37),
				K: toCmyk(0.22),
			},
			wantErr: false,
		}, {
			name: "rgb2cmyk-#2b00ff",
			color: pdf.RGBColor{
				R: uint16(uint64(43 * 0xFFFF / 255)),
				G: 0,
				B: 0xFFFF,
				A: 0xFFFF,
			},
			wantCmyk: pdf.CMYKColor{
				C: toCmyk(0.83),
				M: toCmyk(1.0),
				Y: toCmyk(0.0),
				K: toCmyk(0.0),
			},
			wantErr: false,
		}, {
			name: "cmyk2cmyk",
			color: pdf.CMYKColor{
				C: toCmyk(0.1),
				M: toCmyk(0.2),
				Y: toCmyk(0.3),
				K: toCmyk(0.4),
			},
			wantCmyk: pdf.CMYKColor{
				C: toCmyk(0.1),
				M: toCmyk(0.2),
				Y: toCmyk(0.3),
				K: toCmyk(0.4),
			},
			wantErr: false,
		}, {
			name: "gray2cmyk",
			color: pdf.GrayColor{
				Y: uint16(uint64(55 * 0xFFFF / 255)),
			},
			wantCmyk: pdf.CMYKColor{
				K: toCmyk(0.78),
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

		gray := pdf.GrayColor{
			Y: uint16(uint64(0xCC * 0xFFFF / 255)),
		}

		expectRGB := pdf.RGBColor{
			R: gray.Y,
			G: gray.Y,
			B: gray.Y,
			A: 0xFFFF,
		}
		expectCMYK := pdf.CMYKColor{
			K: toCmyk(0.2),
		}

		rgb, err := pdf.ConvertToRGB(gray)
		assert.NoError(t, err)
		assert.Equal(t, expectRGB, rgb)

		cmyk, err := pdf.ConvertToCMYK(rgb)
		assert.NoError(t, err)
		assert.InDeltaMapValues(t, mapCmyk(expectCMYK), mapCmyk(cmyk), 0.01)
	})
}

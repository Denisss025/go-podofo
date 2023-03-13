package pdf_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/denisss025/go-podofo/internal/pdf"
)

func TestCheckEOL(t *testing.T) {
	t.Parallel()

	args := func(s string) [2]byte {
		return [2]byte{s[0], s[1]}
	}

	tests := []struct {
		name string
		args [2]byte
		want bool
	}{{
		name: "CRLF",
		args: args("\r\n"),
		want: true,
	}, {
		name: "LFCR",
		args: args("\n\r"),
		want: true,
	}, {
		name: "space-CR",
		args: args(" \r"),
		want: true,
	}, {
		name: "space-LF",
		args: args(" \n"),
		want: true,
	}, {
		name: "LFLF",
		args: args("\n\n"),
		want: false,
	}, {
		name: "CRCR",
		args: args("\r\r"),
		want: false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, pdf.CheckEOL(tt.args[0], tt.args[1]))
		})
	}
}

func TestLastIndex(t *testing.T) {
	t.Parallel()

	t.Run("small buffer", func(t *testing.T) {
		t.Parallel()

		const (
			search     = "TEST"
			testString = "12345TEST012345TEST012"
		)

		buf := strings.NewReader(testString)

		i, err := pdf.IndexLast(buf, []byte(search))
		assert.NoError(t, err)
		assert.Equal(t, strings.LastIndex(testString, search), int(i))

		// No error on repeat
		i, err = pdf.IndexLast(buf, []byte(search))
		assert.NoError(t, err)
		assert.Equal(t, strings.LastIndex(testString, search), int(i))

		i, err = pdf.IndexLast(buf, []byte("XREF"))
		assert.NoError(t, err)
		assert.Equal(t, int64(-1), i)
	})

	// TODO: more tests!!!
}

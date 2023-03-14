package pdf_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/denisss025/go-podofo/internal/pdf"
)

var errMockFail = fmt.Errorf("mock failure")

type mockReader struct {
	reader io.ReadSeeker

	failOnSeekStart bool
	failOnSeekEnd   bool
	failOnRead      bool
}

func (r *mockReader) Read(p []byte) (n int, err error) {
	if r.failOnRead {
		return 0, fmt.Errorf("on read: %w", errMockFail)
	}

	return r.reader.Read(p)
}

func (r *mockReader) Seek(offset int64, whence int) (int64, error) {
	switch {
	case whence == io.SeekStart && r.failOnSeekStart:
		return 0, fmt.Errorf("seek start: %w", errMockFail)
	case whence == io.SeekEnd && r.failOnSeekEnd:
		return 0, fmt.Errorf("seek end: %w", errMockFail)
	}

	return r.reader.Seek(offset, whence)
}

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

	const search = "TEST"

	t.Run("small buffer", func(t *testing.T) {
		t.Parallel()

		const testString = "12345TEST012345TEST012"

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

	t.Run("bigger buffer", func(t *testing.T) {
		t.Parallel()

		const n int64 = 10

		x := bytes.Repeat([]byte("This is a test"), 512)

		copy(x[n:], []byte(search))

		buf := bytes.NewReader(x)

		i, err := pdf.IndexLast(buf, []byte(search))
		assert.NoError(t, err)
		assert.Equal(t, n, i)
	})

	t.Run("failures", func(t *testing.T) {
		t.Parallel()

		x := bytes.Repeat([]byte(search), 10)
		buf := bytes.NewReader(x)

		r := &mockReader{reader: buf}

		_, err := pdf.IndexLast(r, []byte(search))
		assert.NoError(t, err)

		r.failOnSeekStart = true
		_, err = pdf.IndexLast(r, []byte(search))
		assert.EqualError(t, err, "last index: seek start: mock failure")

		r.failOnSeekStart = false
		r.failOnSeekEnd = true
		_, err = pdf.IndexLast(r, []byte(search))
		assert.EqualError(t, err, "last index: seek end: mock failure")

		r.failOnSeekEnd = false
		r.failOnRead = true
		_, err = pdf.IndexLast(r, []byte(search))
		assert.EqualError(t, err, "last index: on read: mock failure")
	})
}

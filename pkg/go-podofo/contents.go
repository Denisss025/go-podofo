package podofo

import (
	"bytes"
	"fmt"
	"io"
)

type Contents struct {
	parent Page
	object Object
}

func NewContents(parent *Page, obj Object) (*Contents, error) {
	contents := &Contents{parent: *parent, object: obj}

	if obj == nil {
		contents.object = &Array{
			Objects: parent.Document().Objects,
		}

		if err := contents.reset(); err != nil {
			return nil, fmt.Errorf("new contents: %w", err)
		}
	}

	return contents, nil
}

func (c *Contents) Object() Object { return c.object }

func (c *Contents) WriteTo(w io.Writer) (n int64, err error) {
	// TODO? what about GetStream()?
	if arr, ok := c.object.(*Array); ok {
		return c.writeTo(w, arr)
	}

	dict, isDict := c.object.(*Dictionary)
	if !isDict {
		return 0, fmt.Errorf("write contents: %w", ErrInvalidDataType)
	}

	if n, err = dict.WriteTo(w); err != nil {
		err = fmt.Errorf("write contents: %w", err)
	}

	return n, err
}

func (c *Contents) writeTo(w io.Writer, arr *Array) (n int64, err error) {
	for _, writeObj := range arr.Objects {
		if writeObj == nil {
			continue
		}

		c, err := writeObj.WriteTo(w)
		n += c

		if err != nil {
			return n, fmt.Errorf("contents: write to: %w", err)
		}
	}

	return n, nil
}

func (c *Contents) MarshalBinary() ([]byte, error) {
	buf := &bytes.Buffer{}

	if _, err := c.WriteTo(buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *Contents) reset() error {
	return c.parent.Object().(*Dictionary).AddKeyIndirect(KeyContents, c.object)
}

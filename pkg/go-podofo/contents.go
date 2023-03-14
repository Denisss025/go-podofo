package podofo

import (
	"fmt"

	"github.com/denisss025/go-podofo/internal/pdf"
)

type Contents struct {
	parent Page
	object Object
}

func NewContents(parent *Page, obj Object) (*Contents, error) {
	contents := &Contents{parent: *parent, object: obj}

	if obj == nil {
		contents.object = &Array{
			objects: parent.Document().Objects,
		}

		if err := contents.reset(); err != nil {
			return nil, fmt.Errorf("new contents: %w", err)
		}
	}

	return contents, nil
}

func (c *Contents) Object() Object { return c.object }

func (c *Contents) MarshalPDF(w *pdf.Writer) (err error) {
	// TODO? what about GetStream()?
	if arr, ok := c.object.(*Array); ok {
		return c.writeTo(w, arr)
	}

	dict, isDict := c.object.(*Dictionary)
	if !isDict {
		return fmt.Errorf("write contents: %w", ErrInvalidDataType)
	}

	if err = dict.MarshalPDF(w); err != nil {
		err = fmt.Errorf("write contents: %w", err)
	}

	return err
}

func (c *Contents) writeTo(w *pdf.Writer, arr *Array) (err error) {
	for _, writeObj := range arr.objects {
		if writeObj == nil {
			continue
		}

		if err = writeObj.MarshalPDF(w); err != nil {
			return fmt.Errorf("contents: write to: %w", err)
		}
	}

	return nil
}

// TODO? do we need this?
// func (c *Contents) MarshalBinary() ([]byte, error) {
// 	buf := &bytes.Buffer{}

// 	if _, err := c.WriteTo(buf); err != nil {
// 		return nil, err
// 	}

// 	return buf.Bytes(), nil
// }

func (c *Contents) reset() error {
	return c.parent.Object().(*Dictionary).AddKeyIndirect(KeyContents, c.object)
}

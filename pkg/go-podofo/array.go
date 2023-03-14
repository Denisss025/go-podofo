package podofo

import (
	"errors"
	"fmt"

	"github.com/denisss025/go-podofo/internal/pdf"
)

type Array struct {
	pdf.BaseObject

	cont DataContainer

	objects []Object
}

func (array *Array) Kind() ObjectKind { return pdf.ObjectKindArray }

func (array *Array) MarshalPDF(w *pdf.Writer) (err error) {
	const (
		arrayStart = "["
		arrayEnd   = "]"
		space      = " "
		lf         = "\n"
	)

	// TODO? write mode?

	defer func() {
		if err != nil {
			err = fmt.Errorf("array: write: %w", err)
		}
	}()

	err = errors.Join(
		writeString(w, arrayStart),
		writeStringClean(w, space),
	)
	if err != nil {
		return err
	}

	for i, obj := range array.objects {
		err = obj.MarshalPDF(w)

		if i%10 == 0 {
			err = errors.Join(err, writeStringClean(w, lf))
		}

		if err != nil {
			return err
		}
	}

	return writeString(w, arrayEnd)
}

func (array *Array) NumObjects() int {
	return len(array.objects)
}

func (array *Array) Remove(index int) error {
	if index >= len(array.objects) {
		return fmt.Errorf("remove from array: %w", ErrValueOutOfRange)
	}

	array.objects = append(array.objects[:index], array.objects[index+1:]...)

	array.cont.SetDirty()

	return nil
}

func (array *Array) At(index int) Object {
	obj := array.objects[index]

	if ref, ok := obj.(*Reference); ok {
		return array.cont.GetIndirectObject(ref)
	}

	return obj
}

func (array *Array) Append(objects ...Object) error {
	if cap(array.objects) < len(array.objects)+len(objects) {
		objs := make([]Object, len(array.objects),
			3*(len(array.objects)+len(objects))/2)
		copy(objs, array.objects)

		array.objects = objs
	}

	for _, obj := range objects {
		obj, err := obj.Copy()
		if err != nil {
			return fmt.Errorf("array: append: %w", err)
		}
		obj.SetParent(array)
		array.objects = append(array.objects, obj)
	}

	array.cont.SetDirty()

	return nil
}

func (array *Array) AppendIndirect(objects ...Object) error {
	for _, obj := range objects {
		if !array.cont.IsIndirectReferenceAllowed(obj) {
			return fmt.Errorf("array: append indirect: %w", ErrInvalidHandle)
		}
	}

	if cap(array.objects) < len(array.objects)+len(objects) {
		objs := make([]Object, len(array.objects),
			3*(len(array.objects)+len(objects))/2)
		copy(objs, array.objects)

		array.objects = objs
	}

	for _, obj := range objects {
		array.objects = append(array.objects, obj.GetIndirectReference())
	}

	array.cont.SetDirty()

	return nil
}

func (array *Array) AppendIndirectSafe(objects ...Object) error {
	if cap(array.objects) < len(array.objects)+len(objects) {
		objs := make([]Object, len(array.objects),
			3*(len(array.objects)+len(objects))/2)
		copy(objs, array.objects)

		array.objects = objs
	}

	for _, obj := range objects {
		if array.cont.IsIndirectReferenceAllowed(obj) {
			array.objects = append(array.objects, obj.GetIndirectReference())
		} else {
			cpy, err := obj.Copy()
			if err != nil {
				return fmt.Errorf("array: append indirect safe: %w", err)
			}

			array.objects = append(array.objects, cpy)
		}
	}

	array.cont.SetDirty()

	return nil
}

func (array *Array) Set(index int, obj Object) {
	array.objects[index] = obj
}

func (array *Array) SetIndirect(index int, obj Object) (Object, error) {
	panic("not implemented") // TODO: implement me
}

func (array *Array) SetIndirectSafe(index int, obj Object) (Object, error) {
	panic("not implemented") // TODO: implement me
}

func (array *Array) Clear() {
	if len(array.objects) == 0 {
		return
	}

	for i := range array.objects {
		array.objects[i].SetParent(nil)
	}

	array.cont.SetDirty()
}

func (array *Array) Copy() (obj Object, err error) {
	cpy := &Array{
		cont:    array.cont,
		objects: make([]pdf.Object, 0, len(array.objects)),
	}

	cpy.BaseObject, err = cpy.BaseObject.Copy()
	if err != nil {
		return nil, fmt.Errorf("copy array: %w", err)
	}

	// TODO? is it ok?
	cpy.Append(array.objects...)

	return cpy, err
}

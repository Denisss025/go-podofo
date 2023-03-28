package pdf

import (
	"strconv"
	"strings"
)

type Number struct {
	value string
}

// MarshalPDF encodes the receiver a PDF bytes.
func (num *Number) MarshalPDF(_ *Writer) error {
	panic("not implemented") // TODO: Implement
}

func (num *Number) SetParent(_ Object) {
	panic("not implemented") // TODO: Implement
}

func (num *Number) Parent() Object {
	panic("not implemented") // TODO: Implement
}

func (num *Number) Kind() ObjectKind {
	panic("not implemented") // TODO: Implement
}

func (num *Number) Copy() (Object, error) {
	panic("not implemented") // TODO: Implement
}

func (num *Number) Document() *Document {
	panic("not implemented") // TODO: Implement
}

func (num *Number) Dictionary() *Dictionary {
	panic("not implemented") // TODO: Implement
}

func (num *Number) GetIndirectReference() *Reference {
	panic("not implemented") // TODO: Implement
}

func (num Number) Int() int { return int(num.Int64()) }

func (num Number) Int64() int64 {
	if strings.Contains(num.value, ".") {
		return int64(num.Float64())
	}

	result, _ := strconv.ParseInt(num.value, 10, 64)

	return result
}

func (num Number) Float64() float64 {
	result, _ := strconv.ParseFloat(num.value, 64)

	return result
}

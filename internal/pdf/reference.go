package pdf

import "fmt"

type Generation uint16

const FirstGeneration Generation = 1

type Reference struct {
	ObjectNo     uint32
	GenerationNo Generation
}

var _ = Marshaler((*Reference)(nil))

func (ref Reference) String() string {
	return fmt.Sprintf("%d %d R", ref.ObjectNo, ref.GenerationNo)
}

func (ref Reference) IsIndirect() bool {
	return ref.ObjectNo != 0 || ref.GenerationNo != 0
}

func (ref Reference) MarshalPDF(w *Writer) error {
	const format = "%d %d R"

	if w.HasFlag(WriteFlagNoInlineLiteral) {
		if err := w.WriteByte(' '); err != nil {
			return fmt.Errorf("marshal reference: %w", err)
		}
	}

	_, err := fmt.Fprintf(w, format, ref.ObjectNo, ref.GenerationNo)
	if err != nil {
		err = fmt.Errorf("marshal reference: %w", err)
	}

	return err
}

// func (ref Reference) IsLessThan(other Reference) bool {
// 	return ref.ObjectNo < other.ObjectNo ||
// 		(ref.ObjectNo == other.ObjectNo &&
// 			ref.GenerationNo < other.GenerationNo)
// }

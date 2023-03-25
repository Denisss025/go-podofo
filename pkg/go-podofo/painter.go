package podofo

type Painter struct{}

type TextStyle struct{}

func NewPainter(canvas any) *Painter {
	_ = canvas
	panic("not implemented") // TODO: implement me
}

func (p *Painter) DrawText(text string, x float64, y float64, style ...TextStyle) error {
	panic("not implemented") // TODO: implement me
}

func (p *Painter) FinishDrawing() error {
	panic("not implemented") // TODO: implement me
}

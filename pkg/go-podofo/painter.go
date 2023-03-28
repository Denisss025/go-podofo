package podofo

type painterFlag uint8

const (
	painterFlagNone    painterFlag = 0
	painterFlagPrepend painterFlag = 1 << (iota - 1)
	painterFlagNoSaveRestorePrior
	painterFlagNoSaveRestore
	painterFlagRawCoordinates
)

type PainterFlag func(*painterFlag)

// PainterFlagPrepend does nothing for now.
func PainterFlagPrepend(p *painterFlag) { *p += painterFlagPrepend }

// PainterFlagNoSaveRestorePrior disables Save/Restore or previous content.
// Implies PainterFlagRawCoordinates.
func PainterFlagNoSaveRestorePrior(p *painterFlag) {
	*p += painterFlagNoSaveRestorePrior
}

// PainterFlagNoSaveRestore disables Save/Restore of added content in this
// painting session.
func PainterFlagNoSaveRestore(p *painterFlag) { *p += painterFlagNoSaveRestore }

// PainterFlagRawCoordinates does nothing for now.
func PainterFlagRawCoordinates(p *painterFlag) {
	*p += painterFlagRawCoordinates
}

type painterStatus uint8

const (
	painterStatusDefault painterStatus = 1 << iota
	painterStatusTextObject
	painterStatusTextArray
	painterStatusExtention
)

type Painter struct {
	canvas     CanvasObject
	stackCount int
	tabWidth   uint8
	flags      painterFlag
	status     painterStatus
}

type TextStyle struct{}

func NewPainter(flags ...PainterFlag) *Painter {
	p := &Painter{
		status:   painterStatusDefault,
		tabWidth: 4,
	}

	// TODO: GraphicsState
	// TODO: TextState
	// TODO: TextObject
	// TODO: objStream

	for _, flag := range flags {
		flag(&p.flags)
	}

	return p
}

func (p *Painter) SetCanvas(canvas CanvasObject) {
	if canvas == p.canvas {
		return
	}

	_ = p.finishDrawing() // TODO? ignore error?
	p.reset()
	// TODO: objStream
	p.canvas = canvas
}

func (p *Painter) DrawText(text string, x float64, y float64, style ...TextStyle) error {
	panic("not implemented") // TODO: implement me
}

func (p *Painter) FinishDrawing() error {
	err := p.finishDrawing()
	p.reset()

	return err
}

func (p *Painter) finishDrawing() error {
	panic("not implemented") // TODO: implement me
}

func (p *Painter) reset() {

	panic("not implemented") // TODO: implement me
}

func (p *Painter) Close() error { return p.finishDrawing() }

package podofo

type Font struct{}

type Metrics struct{}

func (f *Font) Metrics() *Metrics {
	panic("not implemented") // TODO: implement me
}

func (m *Metrics) FontName() string {
	panic("not implemented") // TODO: implement me
}

func (m *Metrics) FontFamilyName() string {
	panic("not implemented") // TODO: implement me
}

func (m *Metrics) FilePath() string {
	panic("not implemented") // TODO: implement me
}

func (m *Metrics) FaceIndex() string {
	panic("not implemented") // TODO: implement me
}

package podofo

type Pos struct {
	X float64
	Y float64
}

type Size struct {
	Width  float64
	Height float64
}

type Rect struct {
	Pos
	Size
}

type PageSize func() Rect

func PageSizeA0() Rect {
	return Rect{Size: Size{Width: 2384, Height: 3370}}
}

func PageSizeA1() Rect {
	return Rect{Size: Size{Width: 1684, Height: 2384}}
}

func PageSizeA2() Rect {
	return Rect{Size: Size{Width: 1191, Height: 1684}}
}

func PageSizeA3() Rect {
	return Rect{Size: Size{Width: 842, Height: 1190}}
}

func PageSizeA4() Rect {
	return Rect{Size: Size{Width: 595, Height: 842}}
}

func PageSizeA5() Rect {
	return Rect{Size: Size{Width: 420, Height: 595}}
}

func PageSizeA6() Rect {
	return Rect{Size: Size{Width: 297, Height: 420}}
}

func PageSizeLetter() Rect {
	return Rect{Size: Size{Width: 612, Height: 792}}
}

func PageSizeLegal() Rect {
	return Rect{Size: Size{Width: 612, Height: 1008}}
}

func PageSizeTabloid() Rect {
	return Rect{Size: Size{Width: 792, Height: 1224}}
}

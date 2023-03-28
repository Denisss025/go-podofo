package podofo

import "github.com/denisss025/go-podofo/internal/pdf"

type Rect = pdf.Rect
type PageSize = pdf.PageSize

func PageSizeA0() PageSize      { return pdf.PageSizeFunc(pdf.PageSizeA0) }
func PageSizeA1() PageSize      { return pdf.PageSizeFunc(pdf.PageSizeA1) }
func PageSizeA2() PageSize      { return pdf.PageSizeFunc(pdf.PageSizeA2) }
func PageSizeA3() PageSize      { return pdf.PageSizeFunc(pdf.PageSizeA3) }
func PageSizeA4() PageSize      { return pdf.PageSizeFunc(pdf.PageSizeA4) }
func PageSizeA5() PageSize      { return pdf.PageSizeFunc(pdf.PageSizeA5) }
func PageSizeA6() PageSize      { return pdf.PageSizeFunc(pdf.PageSizeA6) }
func PageSizeLetter() PageSize  { return pdf.PageSizeFunc(pdf.PageSizeLetter) }
func PageSizeLegal() PageSize   { return pdf.PageSizeFunc(pdf.PageSizeLegal) }
func PageSizeTabloid() PageSize { return pdf.PageSizeFunc(pdf.PageSizeTabloid) }

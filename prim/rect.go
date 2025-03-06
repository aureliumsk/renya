package prim

import (
	"github.com/gdamore/tcell/v2"
	"image"
	"renya/cellarray"
)

type Rectangle struct {
	Bounds image.Rectangle
	Style  tcell.Style
}

func (r Rectangle) Display(array cellarray.CellArray) {
	for y := r.Bounds.Min.Y; y < r.Bounds.Max.Y; y++ {
		for x := r.Bounds.Min.X; x < r.Bounds.Max.X; x++ {
			point := image.Point{X: x, Y: y}
			onZeroY := y == r.Bounds.Min.Y
			onZeroX := x == r.Bounds.Min.X
			onMaxX := x == r.Bounds.Max.X-1
			onMaxY := y == r.Bounds.Max.Y-1
			switch true {
			case onZeroX && onZeroY:
				array.Set(point, tcell.RuneULCorner, r.Style)
			case onMaxX && onZeroY:
				array.Set(point, tcell.RuneURCorner, r.Style)

			case onZeroX && onMaxY:
				array.Set(point, tcell.RuneLLCorner, r.Style)
			case onMaxX && onMaxY:
				array.Set(point, tcell.RuneLRCorner, r.Style)

			case onZeroX || onMaxX:
				array.Set(point, tcell.RuneVLine, r.Style)
			case onZeroY || onMaxY:
				array.Set(point, tcell.RuneS7, r.Style)

			default:
				array.Set(point, ' ', r.Style)
			}
		}
	}
}

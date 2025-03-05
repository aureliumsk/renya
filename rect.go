package main

import (
	"github.com/gdamore/tcell/v2"
	"image"
	"renya/cellarray"
)

func rectangle(array cellarray.CellArray, bounds image.Rectangle, style tcell.Style) {
	// style := style.Background(tcell.ColorDefault)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			point := image.Point{X: x, Y: y}
			switch true {
			case x == bounds.Min.X && y == bounds.Min.Y:
				array.Set(point, tcell.RuneULCorner, style)
			case x == bounds.Max.X-1 && y == bounds.Min.Y:
				array.Set(point, tcell.RuneURCorner, style)

			case x == bounds.Min.X && y == bounds.Max.Y-1:
				array.Set(point, tcell.RuneLLCorner, style)
			case x == bounds.Max.X-1 && y == bounds.Max.Y-1:
				array.Set(point, tcell.RuneLRCorner, style)

			case x == bounds.Min.X || x == bounds.Max.X-1:
				array.Set(point, tcell.RuneVLine, style)
			case y == bounds.Min.Y || y == bounds.Max.Y-1:
				array.Set(point, tcell.RuneS9, style)

			default:
				array.Set(point, ' ', style)
			}
		}
	}
}

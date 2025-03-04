package cellarray

import (
	"github.com/gdamore/tcell/v2"
	"image"
)

type Cell struct {
	Rune  rune
	Style tcell.Style
}

type CellArray struct {
	Size    image.Point
	content []Cell
}

func New(size image.Point) CellArray {
	return CellArray{
		Size:    size,
		content: make([]Cell, size.X*size.Y),
	}
}

func (c CellArray) Set(p image.Point, r rune, s tcell.Style) {
	c.content[p.Y*c.Size.X+p.X].Rune = r
	c.content[p.Y*c.Size.X+p.X].Style = s
}

func (c CellArray) Get(p image.Point) Cell {
	return c.content[p.Y*c.Size.X+p.X]
}

// TODO: Написать метод для изменения размера (ширины и высоты) массива

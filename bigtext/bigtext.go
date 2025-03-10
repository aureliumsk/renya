package bigtext

import (
	_ "embed"
	"github.com/gdamore/tcell/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"renya/cellarray"
	// "math"
)

type BigText struct {
	Size    image.Point
	content []rune
}

func (b BigText) Display(array cellarray.CellArray, style tcell.Style, p image.Point) {
	for dy := 0; dy < b.Size.Y; dy++ {
		for dx := 0; dx < b.Size.X; dx++ {
			char := b.content[dy*b.Size.X+dx]
			if char == 0x0 {
				continue
			}
			array.Set(image.Pt(p.X+dx, p.Y+dy), char, style)
		}
	}
}

func Render(s string, face font.Face) BigText {
	textBounds, _ := font.BoundString(face, s)
	startY := -textBounds.Min.Y
	maxY := startY.Floor() + textBounds.Max.Y.Floor()
	maxX := textBounds.Max.X.Floor()

	dot := fixed.Point26_6{
		X: 0,
		Y: startY,
	}

	drawer := font.Drawer{
		Dot:  dot,
		Face: face,
	}

	bounds := image.Rect(0, 0, maxX, maxY)
	src := image.NewUniform(color.Gray{0xff})
	canvas := image.NewGray(bounds)

	drawer.Src = src
	drawer.Dst = canvas

	drawer.DrawString(s)

	text := make([]rune, maxX*maxY)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			pix := canvas.Pix[y*canvas.Stride+x]
			if pix != 0xff {
				continue
			}
			text[y*maxX+x] = '8'
		}
	}
	return BigText{Size: image.Pt(maxX, maxY), content: text}
}

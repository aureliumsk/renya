package main

import (
	_ "embed"
	"github.com/gdamore/tcell/v2"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	// "math"
)

func parseOrPanic(b []byte) *truetype.Font {
	f, err := truetype.Parse(b)
	if err != nil {
		panic(err)
	}
	return f
}

func fixedIntToInt(i fixed.Int26_6) int {
	return int(i >> 6)
}

const (
	points      = 24.0
	dotsPerInch = 72.0
	imageWidth  = 128.0
	imageHeight = 128.0
)

//go:embed resources/HarreeghPoppedCyrillic.ttf
var rawFont []byte
var parsedFont = parseOrPanic(rawFont)

type giantText struct {
	width   int
	height  int
	content []rune
}

func (g giantText) write(screen tcell.Screen, style tcell.Style, p image.Point) {
	for dy := 0; dy < g.height; dy++ {
		for dx := 0; dx < g.width; dx++ {
			char := g.content[dy*g.width+dx]
			if char == 0 {
				char = ' '
			}
			screen.SetContent(p.X+dx, p.Y+dy, char, nil, style)
		}
	}
}

func render(s string) giantText {
	face := truetype.NewFace(parsedFont, &truetype.Options{
		Size: points,
		DPI:  dotsPerInch,
	})
	bounds := image.Rect(0, 0, imageWidth, imageHeight)
	src := image.NewUniform(color.Alpha{0xff})
	canvas := image.NewAlpha(bounds)

	drawer := font.Drawer{
		Dst:  canvas,
		Src:  src,
		Face: face,
	}

	dot := fixed.Point26_6{
		X: 0,
		Y: fixed.Int26_6(points * float64(dotsPerInch) * (64.0 / 72.0)),
	}

	drawer.Dot = dot
	textBounds, _ := drawer.BoundString(s)
	drawer.DrawString(s)

	maxY := fixedIntToInt(textBounds.Max.Y)
	minY := fixedIntToInt(textBounds.Min.Y)
	maxX := fixedIntToInt(textBounds.Max.X)

	text := make([]rune, maxX*(maxY-minY))
	for y := 0; y < (maxY - minY); y++ {
		for x := 0; x < maxX; x++ {
			pix := canvas.Pix[(y+minY)*canvas.Stride+x]
			if pix != 0xff {
				continue
			}
			text[y*maxX+x] = '8'
		}
	}
	return giantText{width: maxX, height: maxY - minY, content: text}
}

package main

import (
	"bytes"
	"image"
	"image/png"

	"github.com/gdamore/tcell/v2"
	"golang.org/x/image/draw"
)

func drawImage(rawImage []byte, screen tcell.Screen, p image.Point, bounds image.Rectangle) {
	reader := bytes.NewReader(rawImage)
	img, err := png.Decode(reader)
	if err != nil {
		panic(err)
	}
	scaledImg := image.NewRGBA(bounds)
	draw.NearestNeighbor.Scale(scaledImg, bounds, img, img.Bounds(), draw.Src, nil)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			offset := scaledImg.PixOffset(x, y)
			if (scaledImg.Pix[offset+3]) != 0xff { // alpha
				continue
			}

			color := tcell.NewRGBColor(
				int32(scaledImg.Pix[offset]),
				int32(scaledImg.Pix[offset+1]),
				int32(scaledImg.Pix[offset+2]),
			)

			screen.SetContent(p.X+x, p.Y+y, '@', nil, tcell.StyleDefault.Foreground(color))
		}
	}
}

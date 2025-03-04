package main

import (
	"bytes"
	_ "embed"
	"github.com/golang/freetype/truetype"
	"image"
	"image/png"
)

func parseImageOrPanic(b []byte) image.Image {
	img, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	return img
}

func parseFontOrPanic(b []byte) *truetype.Font {
	font, err := truetype.Parse(b)
	if err != nil {
		panic(err)
	}
	return font
}

//go:embed resources/rose.png
var rawRoseImage []byte

//go:embed resources/HarreeghPoppedCyrillic.ttf
var rawFont []byte

var roseImage = parseImageOrPanic(rawRoseImage)
var font = parseFontOrPanic(rawFont)

package resources

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

//go:embed resources/song.mp3
var Song []byte

//go:embed resources/rose-flower.png
var rawRoseFlowerImage []byte

//go:embed resources/rose-tilted.png
var rawRoseTiltedImage []byte

//go:embed resources/rose-tilted-flipped.png
var rawRoseTiltedFlippedImage []byte

//go:embed resources/HarreeghPoppedCyrillic.ttf
var rawFont []byte

var RoseFlowerImage = parseImageOrPanic(rawRoseFlowerImage)
var Font = parseFontOrPanic(rawFont)
var RoseTiltedImage = parseImageOrPanic(rawRoseTiltedImage)
var RoseTiltedFlippedImage = parseImageOrPanic(rawRoseTiltedFlippedImage)

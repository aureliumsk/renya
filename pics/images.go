package pics

import (
	"github.com/gdamore/tcell/v2"
	"golang.org/x/image/draw"
	"image"
	"renya/cellarray"
)

func ResizeImage(img image.Image, bounds image.Rectangle) *image.RGBA {
	scaledImg := image.NewRGBA(bounds)
	draw.NearestNeighbor.Scale(scaledImg, bounds, img, img.Bounds(), draw.Src, nil)
	return scaledImg
}

func DrawImage(array cellarray.CellArray, img image.Image, p image.Point, baseStyle tcell.Style) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var color tcell.Color
			switch img := img.(type) {
			case *image.RGBA:
				offset := img.PixOffset(x, y)
				if (img.Pix[offset+3]) != 0xff { // alpha
					continue
				}

				color = tcell.NewRGBColor(
					int32(img.Pix[offset]),
					int32(img.Pix[offset+1]),
					int32(img.Pix[offset+2]),
				)
			default:
				color = tcell.FromImageColor(img.At(x, y))
			}
			array.Set(image.Pt(p.X+x, p.Y+y), '@', baseStyle.Foreground(color))
		}
	}
}

func FlipImage(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	flippedImg := image.NewRGBA(bounds)
	for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
		for i := bounds.Min.X; i < bounds.Max.X; i++ {
			flippedImg.Set(bounds.Max.X-i, j, img.At(i, j))
		}
	}
	return flippedImg
}

package slides

import (
	"github.com/gdamore/tcell/v2"
	"image"
	"renya/bigtext"
	"renya/cellarray"
	"renya/pics"
	"renya/prim"
	"renya/resources"
)

func DrawSlide1(c cellarray.CellArray) {
	size := c.Size
	textBackgroundColor := tcell.NewHexColor(0xdddddd)

	slideBackground := tcell.StyleDefault.Background(tcell.ColorBlack)
	textBackground := tcell.StyleDefault.Background(textBackgroundColor).Foreground(tcell.ColorBlack)
	c.Fill(' ', slideBackground)

	text1 := bigtext.Render("С 8 марта!", headerFace)

	point := image.Pt(
		size.X/2-text1.Size.X/2,
		size.Y/2-text1.Size.Y/2,
	)

	const dx, dy = -2, -2

	prim.Rectangle{
		Bounds: image.Rectangle{
			Min: point.Add(image.Pt(dx, dy)),
			Max: point.Add(text1.Size).Add(image.Pt(-dx, -dy)),
		},
		Style: tcell.StyleDefault.Background(textBackgroundColor).Foreground(tcell.ColorBlack),
	}.Display(c)

	text1.Display(c, textBackground, point)

	flowerBounds := image.Rect(0, 0, 64, 32)
	tiltedBounds := image.Rect(0, 0, 64, 64)

	smallRoseFlowerImage := pics.ResizeImage(resources.RoseFlowerImage, flowerBounds)
	smallRoseTiltedImage := pics.ResizeImage(resources.RoseTiltedImage, tiltedBounds)
	roseFlippedImage := pics.ResizeImage(resources.RoseTiltedFlippedImage, tiltedBounds)

	pics.DrawImage(
		c,
		smallRoseFlowerImage,
		image.Pt(size.X/2-flowerBounds.Max.X/2, point.Y+text1.Size.Y),
		slideBackground,
	)

	pics.DrawImage(
		c,
		smallRoseTiltedImage,
		image.Pt(point.X-10-tiltedBounds.Max.X, size.Y/2-tiltedBounds.Max.Y/2),
		slideBackground,
	)

	pics.DrawImage(
		c,
		roseFlippedImage,
		image.Pt(point.X+text1.Size.X+10, size.Y/2-tiltedBounds.Max.Y/2),
		slideBackground,
	)
}

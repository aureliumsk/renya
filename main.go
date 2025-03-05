package main

import (
	_ "embed"
	"github.com/gdamore/tcell/v2"
	"github.com/golang/freetype/truetype"
	"image"
	"renya/bigtext"
	"renya/cellarray"
	"time"
)

const padding = 10

const dotsPerInch = 72.0

var headerFace = truetype.NewFace(font, &truetype.Options{
	Size: 24.0,
	DPI:  dotsPerInch,
})

func finalize(s tcell.Screen) {
	err := recover()
	s.Fini()
	if err != nil {
		panic(err)
	}
}

type presentation struct {
	slides []cellarray.CellArray
	cursor int
}

func projectSlide(s tcell.Screen, c cellarray.CellArray) {
	emptyStyle := tcell.StyleDefault.Background(tcell.ColorBlack)
	for x := 0; x < c.Size.X; x++ {
		for y := 0; y < c.Size.Y; y++ {
			s.SetContent(x, y, ' ', nil, emptyStyle)
		}
		s.Show()
		time.Sleep(5 * time.Millisecond)
	}
	for x := 0; x < c.Size.X; x++ {
		for y := 0; y < c.Size.Y; y++ {
			cell := c.Get(image.Pt(x, y))
			s.SetContent(x, y, cell.Rune, nil, cell.Style)
		}
		s.Show()
		time.Sleep(5 * time.Millisecond)
	}
}

func setBackground(c cellarray.CellArray, style tcell.Style) {
	for y := 0; y < c.Size.Y; y++ {
		for x := 0; x < c.Size.X; x++ {
			c.Set(image.Pt(x, y), ' ', style)
		}
	}
}

func main() {
	createLogger()
	defer closeLogger()
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	s.Init()
	defer finalize(s)

	s.Clear()

	width, height := s.Size()
	size := image.Pt(width, height)

	slide1 := cellarray.New(size)
	slide2 := cellarray.New(size)

	slideBackground := tcell.StyleDefault.Background(tcell.ColorBlack)
	textBackground := tcell.StyleDefault.Background(tcell.ColorGray).Foreground(tcell.ColorBlack)
	setBackground(slide1, slideBackground)
	setBackground(slide2, slideBackground)

	text1 := bigtext.Render("С 8 марта!", headerFace)
	text2 := bigtext.Render("(здесь точно что-то будет)", headerFace)

	text1.Display(slide1, textBackground, image.Pt(
		size.X/2-text1.Size.X/2,
		size.Y/2-text1.Size.Y/2,
	))

	text2.Display(slide2, textBackground, image.Pt(
		size.X/2-text2.Size.X/2,
		size.Y/2-text2.Size.Y/2,
	))

	bounds := image.Rect(0, 0, 64, 32)

	rectangle(slide1, bounds, tcell.StyleDefault.Background(tcell.ColorGreen))

	smallRoseImage := resizeImage(roseImage, bounds)

	drawImage(slide1, smallRoseImage, image.Pt(size.X/2-bounds.Max.X/2, size.Y-bounds.Max.Y-padding), slideBackground)

	pres := presentation{slides: []cellarray.CellArray{slide1, slide2}, cursor: 0}
	projectSlide(s, pres.slides[pres.cursor])

	for {
		s.Show()
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			r := ev.Key()
			switch r {
			case tcell.KeyRune:
				return
			case tcell.KeyRight:
				if pres.cursor >= len(pres.slides)-1 {
					break
				}
				pres.cursor++
				projectSlide(s, pres.slides[pres.cursor])
			case tcell.KeyLeft:
				if pres.cursor <= 0 {
					break
				}
				pres.cursor--
				projectSlide(s, pres.slides[pres.cursor])
			}
		}
	}

}

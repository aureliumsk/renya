package main

import (
	_ "embed"
	"github.com/gdamore/tcell/v2"
	"image"
	"renya/cellarray"
	// "time"
	"renya/slides"
)

const padding = 10

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
	// emptyStyle := tcell.StyleDefault.Background(tcell.ColorBlack)
	// for x := 0; x < c.Size.X; x++ {
	// 	for y := 0; y < c.Size.Y; y++ {
	// 		s.SetContent(x, y, ' ', nil, emptyStyle)
	// 	}
	// 	s.Show()
	// 	time.Sleep(5 * time.Millisecond)
	// }
	for x := 0; x < c.Size.X; x++ {
		for y := 0; y < c.Size.Y; y++ {
			cell := c.Get(image.Pt(x, y))
			s.SetContent(x, y, cell.Rune, nil, cell.Style)
		}
		// s.Show()
		// time.Sleep(5 * time.Millisecond)
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

	slides.DrawSlide1(slide1)

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

package main

import (
	_ "embed"
	"github.com/gdamore/tcell/v2"
	"image"
)

//go:embed resources/rose.png
var rose []byte

func main() {
	createLogger()
	defer closeLogger()
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	s.Init()
	s.Clear()
	defer func() {
		e := recover()
		s.Fini()
		if e != nil {
			panic(e)
		}
	}()

	// style := tcell.StyleDefault.Background(tcell.ColorDarkGrey).Foreground(tcell.ColorWhite)
	// text := render("С 8 МАРТА!")

	width, height := s.Size()

	// text.write(s, style, image.Pt((width/2)-(text.width/2), (height/2)-(text.height/2)))

	bounds := image.Rect(0, 0, 64, 64)

	drawImage(rose, s, image.Pt(width/2-bounds.Max.X/2, height/2-bounds.Max.Y/2), bounds)

	for {
		s.Show()
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Rune() == 'q' {
				return
			}
		}
	}

}

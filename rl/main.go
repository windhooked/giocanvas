package main

import (
	"flag"
	"image/color"
	"math/rand"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/unit"
	gc "github.com/ajstarks/giocanvas"
)

func rn(n int) float32 {
	return float32(rand.Intn(n))
}

func rl(title string, w, h, nlines int, thickness float32) {
	width, height := float32(w), float32(h)
	win := app.NewWindow(app.Title(title), app.Size(unit.Dp(width), unit.Dp(height)))
	for e := range win.Events() {
		if e, ok := e.(system.FrameEvent); ok {
			canvas := gc.NewCanvas(width, height, e.Config, e.Queue, e.Size)

			canvas.Background(gc.ColorLookup("black"))
			for i := 0; i < nlines; i++ {
				r := uint8(rand.Intn(230))
				c := color.RGBA{r, r, r, 150}
				canvas.Line(rn(100), 0, rn(100), 100, thickness, c)
			}

			e.Frame(canvas.Context.Ops)
		}
	}
}

func main() {
	var w, h, n int
	var size float64
	flag.IntVar(&w, "width", 1000, "canvas width")
	flag.IntVar(&h, "height", 1000, "canvas height")
	flag.IntVar(&n, "n", 500, "number of lines")
	flag.Float64Var(&size, "size", 2, "line thickness (%)")

	flag.Parse()
	rand.Seed(time.Now().Unix())
	go rl("Random Lines", w, h, n, float32(size))
	app.Main()
}
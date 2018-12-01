package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/gamestate"
)

const (
	winWidth  = 1270
	winHeight = 900
	winTitle  = "Catastrophy"
)

var (
	winBounds = pixel.R(0, 0, winWidth, winHeight)
)

func run() {
	catlog.Debug("Game launched")

	cfg := pixelgl.WindowConfig{
		Title:  winTitle,
		Bounds: winBounds,
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		catlog.Fatalf("Could not create window: %v", err)
	}

	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		gamestate.Update(dt, win)
		gamestate.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

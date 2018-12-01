package main

import (
	"image/color"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/consts"
	"github.com/wayovertheregaming/catastrophy/dialogue"
	"github.com/wayovertheregaming/catastrophy/gamestate"
	"github.com/wayovertheregaming/catastrophy/levels"
)

const (
	winTitle = "Catastrophy"
)

var (
	backgroundColour = color.RGBA{0x00, 0x00, 0x1a, 0x00}
	winBounds        = pixel.R(0, 0, consts.WinWidth, consts.WinHeight)
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

	// Set the initial level
	gamestate.SetLevel(levels.Ground)

	last := time.Now()

	d := []dialogue.Dialogue{
		dialogue.Dialogue{
			IsPlayer: false,
			Name:     "Name",
			Text:     "Hello",
		},
		dialogue.Dialogue{
			IsPlayer: true,
			Text:     "Hello back",
		},
	}

	for !win.Closed() {
		win.Clear(backgroundColour)

		dt := time.Since(last).Seconds()
		last = time.Now()

		dialogue.Update(dt, win)

		// TEST lines
		if win.JustPressed(pixelgl.KeyP) {
			dialogue.Start(d)
		}
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			catlog.Debug(win.MousePosition())
		}

		gamestate.Update(dt, win)
		gamestate.Draw(win)

		// Draw dialogue on top of other layers
		dialogue.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

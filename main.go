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
	"github.com/wayovertheregaming/catastrophy/player"
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

	// gameView is a batch drawing element for all game view things
	// Will allow us to overlay stuff such as HUD and notifications
	gameView := pixelgl.NewCanvas(win.Bounds())

	// Set the initial level
	gamestate.SetLevel(levels.Ground)

	last := time.Now()

	for !win.Closed() {
		win.Clear(backgroundColour)
		gameView.Clear(backgroundColour)

		dt := time.Since(last).Seconds()
		last = time.Now()

		gamestate.Update(dt, win)
		gamestate.Draw(gameView)

		dialogue.Update(dt, win)

		// inverseMoved is the player position inversed
		inverseMoved := player.GetPos().Scaled(-1)
		// This shift is effectively doing camera controls -- woo!
		gameView.Draw(win, pixel.IM.Moved(inverseMoved))

		// Draw dialogue on top of other layers
		dialogue.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

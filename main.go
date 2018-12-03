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
)

func run() {
	catlog.Debug("Game launched")

	cfg := pixelgl.WindowConfig{
		Title:  winTitle,
		Bounds: consts.WinBounds,
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		catlog.Fatalf("Could not create window: %v", err)
	}

	// Set the initial level
	gamestate.SetLevel(levels.Ground)
	// Start unpaused
	gamestate.UnPauseGame()

	last := time.Now()

	for !win.Closed() {
		win.Clear(backgroundColour)
		consts.GameView.Clear(color.Transparent)

		dt := time.Since(last).Seconds()
		last = time.Now()

		gamestate.Update(dt, win)
		gamestate.Draw()

		dialogue.Update(dt, win)

		// Shift the camera for the background
		cam := pixel.IM.Moved(consts.WinBounds.Center().Sub(player.GetPos()))
		consts.GameView.Draw(win, cam)

		// Draw dialogue on top of other layers
		dialogue.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

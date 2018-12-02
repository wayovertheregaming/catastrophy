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

// func test() {
// 	catlog.Debug(userinput.GetUserInput())
// }

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

	// gameView is a batch drawing element for all game view things
	// this restricts the size of the level to 3000x3000
	// TODO(too much work with time restrictions to make this dynamic, but can be
	// done)
	gameView := pixelgl.NewCanvas(pixel.R(0, 0, 3000, 3000))

	// Set the initial level
	gamestate.SetLevel(levels.Ground)

	last := time.Now()

	for !win.Closed() {
		win.Clear(backgroundColour)
		gameView.Clear(backgroundColour)
		consts.ImdLayer.Clear()

		dt := time.Since(last).Seconds()
		last = time.Now()

		gamestate.Update(dt, win)
		gamestate.Draw(gameView)

		dialogue.Update(dt, win)

		// userinput.Update(win)
		// userinput.Draw()

		// Shift the camera for the background
		cam := pixel.IM.Moved(consts.WinBounds.Center().Sub(player.GetPos()))
		gameView.Draw(win, cam)

		if win.JustPressed(pixelgl.KeyP) {
			// go test()
		}

		// Draw ImDraw shape layer
		consts.ImdLayer.Draw(win)
		// Draw dialogue on top of other layers
		dialogue.Draw(win)
		win.Update()
	}
}

func main() {
	catlog.Debug("first")
	pixelgl.Run(run)
}

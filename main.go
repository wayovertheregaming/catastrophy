package main

import (
	"image/color"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/consts"
	"github.com/wayovertheregaming/catastrophy/decorations"
	"github.com/wayovertheregaming/catastrophy/dialogue"
	"github.com/wayovertheregaming/catastrophy/gamestate"
	"github.com/wayovertheregaming/catastrophy/levelchanger"
	"github.com/wayovertheregaming/catastrophy/levels"
	"github.com/wayovertheregaming/catastrophy/player"
	"github.com/wayovertheregaming/catastrophy/util"
	"github.com/wayovertheregaming/catastrophy/util/userinput"
)

const (
	winTitle = "Catastrophy"
	iconPath = "assets/graphics/catSleep1.png"
)

var (
	backgroundColour = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
)

func run() {
	catlog.Debug("Game launched")

	catlog.Debug("Loading window icon")
	windowIcon := util.LoadPic(iconPath)

	cfg := pixelgl.WindowConfig{
		Title:  winTitle,
		Bounds: consts.WinBounds,
		VSync:  true,
		Icon:   []pixel.Picture{windowIcon},
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		catlog.Fatalf("Could not create window: %v", err)
	}

	consts.TextLayer = pixelgl.NewCanvas(consts.WinBounds)
	consts.PlayerLayer = pixelgl.NewCanvas(consts.WinBounds)

	// Set the initial level
	gamestate.SetLevel(levels.Ground)
	// Start unpaused
	gamestate.UnPauseGame()

	last := time.Now()

	// HowToPlay is called at launch, it will tell the player how to play the game
	dialogue.Start(dialogue.HowToPlay)

	for !win.Closed() {
		win.Clear(backgroundColour)
		consts.TextLayer.Clear(color.Transparent)
		consts.ImdLayer.Clear()
		consts.GameView.Clear(color.Transparent)
		consts.PlayerLayer.Clear(color.Transparent)
		consts.DecorationsLayer.Clear()

		if win.JustPressed(pixelgl.KeyP) {
			levelchanger.Sleep()
		}
		if win.JustPressed(pixelgl.KeyO) {
			catlog.Debug(player.GetPos())
		}

		dt := time.Since(last).Seconds()
		last = time.Now()

		gamestate.Update(dt, win)
		gamestate.Draw()

		decorations.Draw()

		dialogue.Update(dt, win)
		dialogue.Draw()

		userinput.Update(win)
		userinput.Draw()

		// Shift the camera for the background
		cam := pixel.IM.Moved(consts.WinBounds.Center().Sub(player.GetPos()))
		consts.GameView.Draw(win, cam)

		// Draw the decorations to window
		consts.DecorationsLayer.Draw(win)

		// Draw the player
		consts.PlayerLayer.Draw(win, pixel.IM.Moved(consts.WinCentre))

		// Draw ImDraw shape layer
		consts.ImdLayer.Draw(win)
		consts.TextLayer.Draw(win, pixel.IM.Moved(consts.WinBounds.Center()))
		// Draw dialogue on top of other layers
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

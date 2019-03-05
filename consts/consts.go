// Package consts holds values used in multiple parts of the application, but
// cannot be passed around
package consts

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	// WinWidth is the pixel width to set the game window to
	WinWidth = 1270
	// WinHeight is the pixel height to set the game window to
	WinHeight = 900

	// PlayerSide is the size of one side of the player box
	PlayerSide = 50
)

// Level names
const (
	LevelNameGround    = "Ground"
	LevelNameFirst     = "First"
	LevelNameShadow    = "Shadow Realm"
	LevelNameBowldOver = "Bowl'd Over"
)

var (
	// WinBounds is the bounds of the viewer window
	WinBounds = pixel.R(0, 0, WinWidth, WinHeight)
	// WinCentre is the centre of the window as a vector
	WinCentre = pixel.V(WinWidth/2, WinHeight/2)

	// PlayerScale is how much to scale the player by
	PlayerScale float64 = 4
	// PlayerSize is the width and height of the player, assuming the player is
	// facing upwards
	PlayerSize = pixel.V(PlayerSide, PlayerSide)

	// ImdLayer is used as the foreground drawing object.  This is drawn to the
	// window late in the draw process so will be above most things
	ImdLayer = imdraw.New(nil)

	// TextLayer is used as a foreground text drawing layer.  This is drawn to the
	// window late in the draw process so will be above most things
	TextLayer *pixelgl.Canvas

	// GameView is a batch drawing element for all game view things. It is set at
	// level init
	GameView *pixelgl.Canvas

	// DecorationsLayer is a batch layer mapped to the decorations spritesheet for
	// efficient drawing
	DecorationsLayer *pixel.Batch

	// PlayerLayer lets us put the player above decorations
	PlayerLayer *pixelgl.Canvas

	// SleepFor is how long the player sleeps when transitioning to/from shadow
	// realm
	SleepFor = time.Second * 2
)

// SpritePic holds a pixel sprite and Picture
type SpritePic struct {
	Sprite *pixel.Sprite
	Pic    *pixel.Picture
}

// Package consts holds values used in multiple parts of the application, but
// cannot be passed around
package consts

import "github.com/faiface/pixel"

const (
	// WinWidth is the pixel width to set the game window to
	WinWidth = 1270
	// WinHeight is the pixel height to set the game window to
	WinHeight = 900

	// PlayerSide is the size of one side of the player box
	PlayerSide = 50
)

var (
	// WinBounds is the bounds of the viewer window
	WinBounds = pixel.R(0, 0, WinWidth, WinHeight)
	// WinCentre is the centre of the window as a vector
	WinCentre = pixel.V(WinWidth/2, WinHeight/2)

	// PlayerSize is the width and height of the player, assuming the player is
	// facing upwards
	// Note, not tested, just placeholder sizes
	PlayerSize = pixel.V(PlayerSide, PlayerSide)
)

// Package consts holds values used in multiple parts of the application, but
// cannot be passed around
package consts

import "github.com/faiface/pixel"

const (
	// WinWidth is the pixel width to set the game window to
	WinWidth = 2500
	// WinHeight is the pixel height to set the game window to
	WinHeight = 2000
)

var (
	// WinBounds is the bounds of the viewer window
	WinBounds = pixel.R(0, 0, WinWidth, WinHeight)
	// WinCentre is the centre of the window as a vector
	WinCentre = pixel.V(WinWidth/2, WinHeight/2)
)

package gamestate

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Leveller represents something that can be used as a level
type Leveller interface {
	Update(float64, *pixelgl.Window)
	Init()
	Draw(pixel.Target)
	Name() string
	Bounds() pixel.Rect
}

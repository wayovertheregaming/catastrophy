package util

import "github.com/faiface/pixel"

// RectCollides works out if two rectangles collide
func RectCollides(r1, r2 pixel.Rect) bool {
	return pixel.R(0, 0, 0, 0) != r1.Intersect(r2)
}

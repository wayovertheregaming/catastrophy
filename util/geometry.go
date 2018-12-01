package util

import "github.com/faiface/pixel"

// RectCollides works out if two rectangles collide
func RectCollides(r1, r2 pixel.Rect) bool {
	for _, v := range RectCorners(r2) {
		if r1.Contains(v) {
			return true
		}
	}
	return false
}

// RectCorners gets each corner of a `pixel.Rect` as a slice of `pixel.Vec`s
func RectCorners(r pixel.Rect) (corners []pixel.Vec) {
	corners = append(corners, pixel.V(r.Min.X, r.Min.Y))
	corners = append(corners, pixel.V(r.Max.X, r.Min.Y))
	corners = append(corners, pixel.V(r.Min.X, r.Max.Y))
	corners = append(corners, pixel.V(r.Max.X, r.Max.Y))

	return corners
}

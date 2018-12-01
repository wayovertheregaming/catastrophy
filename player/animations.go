package player

import (
	"github.com/faiface/pixel"
	"github.com/wayovertheregaming/catastrophy/util"
)

// These are the images for each frame of animation
// They map the file name to the location on the image file
var (
	sleeping = map[string]pixel.Rect{
		"assets/graphics/catSleep1": pixel.R(0, 0, 100, 200),
	}
	idling = map[string]pixel.Rect{
		"assets/graphics/catIdle1": pixel.R(0, 0, 100, 200),
		"assets/graphics/catIdle2": pixel.R(0, 0, 100, 200),
	}
	sitting = map[string]pixel.Rect{
		"assets/graphics/catSit1": pixel.R(0, 0, 100, 200),
		"assets/graphics/catSit2": pixel.R(0, 0, 100, 200),
	}
	walking = map[string]pixel.Rect{
		"assets/graphics/catWalk1": pixel.R(0, 0, 100, 200),
		"assets/graphics/catWalk2": pixel.R(0, 0, 100, 200),
	}
)

var (
	sleepingSprites = make(map[*pixel.Sprite]pixel.Picture)
	idlingSprites   = make(map[*pixel.Sprite]pixel.Picture)
	sittingSprites  = make(map[*pixel.Sprite]pixel.Picture)
	walkingSprites  = make(map[*pixel.Sprite]pixel.Picture)
)

func init() {
	// Get the sprites from the image name list
	for imgPath, r := range sleeping {
		s, p := util.LoadSprite(imgPath, r)
		sleepingSprites[s] = p
	}
}

package player

import (
	"github.com/faiface/pixel"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/util"
)

const (
	// frameRate is the number of seconds each animation frame is displayed for
	frameRate = 0.5
)

// These are the images for each frame of animation
// They map the file name to the location on the image file
var (
	sleeping = map[string]pixel.Rect{
		"assets/graphics/catSleep1.png": pixel.R(0, 0, 100, 200),
	}
	idling = map[string]pixel.Rect{
		"assets/graphics/catIdle1.png": pixel.R(0, 0, 100, 200),
		"assets/graphics/catIdle2.png": pixel.R(0, 0, 100, 200),
	}
	sitting = map[string]pixel.Rect{
		"assets/graphics/catSit1.png": pixel.R(0, 0, 100, 200),
		"assets/graphics/catSit2.png": pixel.R(0, 0, 100, 200),
	}
	walking = map[string]pixel.Rect{
		"assets/graphics/catWalk1.png": pixel.R(0, 0, 100, 200),
		"assets/graphics/catWalk2.png": pixel.R(0, 0, 100, 200),
	}
)

var (
	sleepingSprites = []spritePic{}
	idlingSprites   = []spritePic{}
	sittingSprites  = []spritePic{}
	walkingSprites  = []spritePic{}
)

type spritePic struct {
	sprite *pixel.Sprite
	pic    *pixel.Picture
}

func init() {
	sleepingSprites = loadIntoVar(sleeping)
	idlingSprites = loadIntoVar(idling)
	sittingSprites = loadIntoVar(sitting)
	walkingSprites = loadIntoVar(walking)
}

func loadIntoVar(assetList map[string]pixel.Rect) []spritePic {
	retSlice := []spritePic{}

	// Get the sprites from the image name list
	for imgPath, r := range assetList {
		catlog.Debugf("Loading player animation sprite %s", imgPath)

		s, p := util.LoadSprite(imgPath, r)
		sp := spritePic{s, &p}
		retSlice = append(retSlice, sp)
	}

	return retSlice
}

// getAnimationFrame returns the sprite/picture of the sprite to display
func getAnimationFrame() {

}

func stateFrameToSprites(state, frame int) spritePic {
	return stateToSprites(state)[frame]
}

func stateToSprites(state int) []spritePic {
	switch state {
	case animationStateSleep:
		return sleepingSprites
	case animationStateIdle:
		return idlingSprites
	case animationStateSitting:
		return sittingSprites
	case animationStateWalking:
		return walkingSprites
	default:
		catlog.Infof("Could not find animation state with index %d, returning idle", state)
		return idlingSprites
	}
}

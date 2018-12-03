package player

import (
	"github.com/faiface/pixel"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/consts"
	"github.com/wayovertheregaming/catastrophy/util"
)

const (
	// frameRate is the number of seconds each animation frame is displayed for
	frameRate = 0.15
)

// These are the images for each frame of animation
// They map the file name to the location on the image file
var (
	sleeping = map[string]pixel.Rect{
		"assets/graphics/catSleep1.png": pixel.R(0, 0, consts.PlayerSide, consts.PlayerSide),
		"assets/graphics/catSleep2.png": pixel.R(0, 0, consts.PlayerSide, consts.PlayerSide),
	}
	idling = map[string]pixel.Rect{
		"assets/graphics/catIdle1.png": pixel.R(0, 0, consts.PlayerSide, consts.PlayerSide),
		"assets/graphics/catIdle2.png": pixel.R(0, 0, consts.PlayerSide, consts.PlayerSide),
		"assets/graphics/catIdle3.png": pixel.R(0, 0, consts.PlayerSide, consts.PlayerSide),
		"assets/graphics/catIdle4.png": pixel.R(0, 0, consts.PlayerSide, consts.PlayerSide),
	}
	sitting = map[string]pixel.Rect{
		"assets/graphics/catSit1.png": pixel.R(0, 0, consts.PlayerSide, consts.PlayerSide),
		"assets/graphics/catSit2.png": pixel.R(0, 0, consts.PlayerSide, consts.PlayerSide),
	}
	walking = map[string]pixel.Rect{
		"assets/graphics/catWalk1.png": pixel.R(0, 0, consts.PlayerSide, consts.PlayerSide),
		"assets/graphics/catWalk2.png": pixel.R(0, 0, consts.PlayerSide, consts.PlayerSide),
		"assets/graphics/catWalk3.png": pixel.R(0, 0, consts.PlayerSide, consts.PlayerSide),
		"assets/graphics/catWalk4.png": pixel.R(0, 0, consts.PlayerSide, consts.PlayerSide),
	}
)

var (
	sleepingSprites = []consts.SpritePic{}
	idlingSprites   = []consts.SpritePic{}
	sittingSprites  = []consts.SpritePic{}
	walkingSprites  = []consts.SpritePic{}
)

func init() {
	catlog.Debug("Doing animation init")

	sleepingSprites = loadIntoVar(sleeping)
	idlingSprites = loadIntoVar(idling)
	sittingSprites = loadIntoVar(sitting)
	walkingSprites = loadIntoVar(walking)
}

func loadIntoVar(assetList map[string]pixel.Rect) []consts.SpritePic {
	retSlice := []consts.SpritePic{}

	// Get the sprites from the image name list
	for imgPath, r := range assetList {
		catlog.Debugf("Loading player animation sprite %s", imgPath)

		s, p := util.LoadSprite(imgPath, r)
		sp := consts.SpritePic{
			Sprite: s,
			Pic:    &p,
		}
		retSlice = append(retSlice, sp)
	}

	return retSlice
}

func stateFrameToSprites(state, frame int) consts.SpritePic {
	return stateToSprites(state)[frame]
}

func stateToSprites(state int) []consts.SpritePic {
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

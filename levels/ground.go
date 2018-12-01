package levels

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/util"
)

const (
	groundImagePath = "assets/graphics/groundFloor.png"
)

var (
	// Ground is the ground level
	Ground = Level{
		name:       "Ground",
		updateFunc: updateGround,
		drawFunc:   drawGround,
		initFunc:   initGround,
	}

	backgroundSprite *pixel.Sprite
	backgroundPic    pixel.Picture

	// groundImageDimensions is effectively the size of the image
	groundImageDimensions = pixel.R(0, 0, 1000, 1000)
)

func init() {
	catlog.Debug("Preparing ground level")

	backgroundSprite, backgroundPic = util.LoadSprite(groundImagePath, groundImageDimensions)
}

func initGround() {

}

func updateGround(dt float64, win *pixelgl.Window) {

}

func drawGround(target pixel.Target) {

}

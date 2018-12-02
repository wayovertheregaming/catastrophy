package levels

import (
	"image"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/player"
	"github.com/wayovertheregaming/catastrophy/util"
)

const (
	shadowImagePath = "assets/graphics/shadowRealm.png"
)

var (
	// ShadowRealm is the ground level
	ShadowRealm = &Level{
		name:          "Shadow Realm",
		updateFunc:    updateShadow,
		drawFunc:      drawShadow,
		initFunc:      initShadow,
		displayPlayer: true,
	}

	shadowBackgroundSprite *pixel.Sprite
	shadowBackgroundPic    pixel.Picture

	// shadowImageDimensions is effectively the size of the image
	shadowImageDimensions pixel.Rect
	shadowFloorStartPos   = pixel.V(500, 200)
)

func init() {
	catlog.Debug("Preparing shadow realm level")

	// Load the background image
	shadowBackgroundSprite, shadowBackgroundPic = util.LoadSprite(shadowImagePath, shadowImageDimensions)

	// shadowImageConfig returns dimensions of groundImagePath
	shadowImageConfig, _, err := image.DecodeConfig(util.GetReaderFromFile(groundImagePath))
	if err != nil {
		catlog.Fatalf("Could not load shadow realm image %v", err)
	}

	// shadowmageDimensions is effectively the size of the image
	shadowImageDimensions = pixel.R(0, 0, float64(shadowImageConfig.Width), float64(shadowImageConfig.Height))
	ShadowRealm.bounds = shadowImageDimensions

	// Load the background image
	shadowBackgroundSprite, shadowBackgroundPic = util.LoadSprite(shadowImagePath, shadowImageDimensions)
}

func initShadow() {
	player.SetPos(shadowFloorStartPos)
}

func updateShadow(dt float64, win *pixelgl.Window) {
	// Try move the player
	if !movePlayer(win, dt, groundFloorCollisions) {
		// Player is not moving, update animation
		player.AnimateIdle()
	}
}

func drawShadow(target pixel.Target) {
	// inverseMoved is the player position inversed
	shadowBackgroundSprite.Draw(target, pixel.IM)
}

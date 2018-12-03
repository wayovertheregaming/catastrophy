package levels

import (
	"image"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/consts"
	"github.com/wayovertheregaming/catastrophy/player"
	"github.com/wayovertheregaming/catastrophy/util"
)

const (
	shadowImagePath = "assets/graphics/shadowRealm.png"
	catGodImagePath = "assets/graphics/catGod.png"
)

var (
	// ShadowRealm is the ground level
	ShadowRealm = &Level{
		name:          consts.LevelNameShadow,
		updateFunc:    updateShadow,
		drawFunc:      drawShadow,
		initFunc:      initShadow,
		displayPlayer: true,
	}

	shadowBackgroundSprite *pixel.Sprite
	shadowBackgroundPic    pixel.Picture

	// shadowImageDimensions is effectively the size of the image
	shadowImageDimensions pixel.Rect
	shadowFloorStartPos   = pixel.V(40, -1000)

	catGodSprite       *pixel.Sprite
	catGodShift        = pixel.ZV
	catGodShiftCounter float64
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

	// Get cat god
	catGodImageConfig, _, err := image.DecodeConfig(util.GetReaderFromFile(catGodImagePath))
	if err != nil {
		catlog.Fatalf("Could not load cat god image %v", err)
	}

	catGodImageDimensions := pixel.R(0, 0, float64(catGodImageConfig.Width), float64(catGodImageConfig.Height))
	catGodSprite, _ = util.LoadSprite(catGodImagePath, catGodImageDimensions)
}

func initShadow() {
	player.SetPos(shadowFloorStartPos)
}

func updateShadow(dt float64, win *pixelgl.Window) {
	// Try move the player
	if !movePlayer(win, dt, []pixel.Rect{}) {
		// Player is not moving, update animation
		player.AnimateIdle()
	}

	// Move cat god up and down
	catGodShiftCounter += dt
	catGodShift.Y = math.Sin(catGodShiftCounter) * 10
}

func drawShadow() {
	// inverseMoved is the player position inversed
	shadowBackgroundSprite.Draw(consts.GameView, pixel.IM.Moved(shadowImageDimensions.Center()))

	catGodSprite.Draw(
		consts.GameView,
		pixel.IM.Moved(shadowImageDimensions.Center().Add(catGodShift)).Scaled(shadowImageDimensions.Center(), 3),
	)
}

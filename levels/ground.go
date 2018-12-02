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
	groundImagePath     = "assets/graphics/groundFloor.png"
	groundCollisionPath = "assets/csv/groundFloorWalls.csv"
)

var (
	// Ground is the ground level
	Ground = &Level{
		name:          "Ground",
		updateFunc:    updateGround,
		drawFunc:      drawGround,
		initFunc:      initGround,
		displayPlayer: true,
	}

	groundBackgroundSprite *pixel.Sprite
	groundBackgroundPic    pixel.Picture

	// groundFloorCollisions are all the rectangles which should cause the player
	// to collide: i.e. unpassable
	groundFloorCollisions []pixel.Rect
	groundImageDimensions pixel.Rect
	groundFloorStartPos   = pixel.V(0, 0)
)

func init() {
	catlog.Debug("Preparing ground level")

	// Get all collision bounds from the CSV file
	groundFloorCollisions = loadCollisions(groundCollisionPath)

	//groundImageConfig returns dimensions of groundImagePath
	groundImageConfig, _, err := image.DecodeConfig(util.GetReaderFromFile(groundImagePath))
	if err != nil {
		catlog.Fatalf("Could not load ground image %v", err)
	}

	// groundImageDimensions is effectively the size of the image
	groundImageDimensions = pixel.R(0, 0, float64(groundImageConfig.Width), float64(groundImageConfig.Height))

	// Load the background image
	groundBackgroundSprite, groundBackgroundPic = util.LoadSprite(groundImagePath, groundImageDimensions)
}

func initGround() {
	player.SetPos(groundFloorStartPos)
}

func updateGround(dt float64, win *pixelgl.Window) {
	// Try move the player
	if !movePlayer(win, dt, groundFloorCollisions) {
		// Player is not moving, update animation
		player.AnimateIdle()
	}
}

func drawGround(target pixel.Target) {
	groundBackgroundSprite.Draw(target, pixel.IM.Moved(groundImageDimensions.Center()))
}

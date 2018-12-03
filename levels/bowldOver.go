package levels

import (
	"image"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/consts"
	"github.com/wayovertheregaming/catastrophy/player"
	"github.com/wayovertheregaming/catastrophy/util"
)

const (
	bowldOverImagePath     = "assets/graphics/bowldOver.png"
	bowldOverCollisionPath = "assets/csv/bowldOverWalls.csv"
)

var (
	// BowldOver is a minigame
	BowldOver = &Level{
		name:          consts.LevelNameBowldOver,
		updateFunc:    updateBowldOver,
		drawFunc:      drawBowldOver,
		initFunc:      initBowldOver,
		displayPlayer: true,
	}

	bowldOverBackgroundSprite *pixel.Sprite
	bowldOverBackgroundPic    pixel.Picture

	// bowldOverCollisions are all the rectangles which should cause the player
	// to collide: i.e. unpassable
	bowldOverCollisions      []pixel.Rect
	bowldOverImageDimensions pixel.Rect
	bowldOverStartPos        = pixel.V(0, 40)
)

func init() {
	catlog.Debug("Preparing bowldOver")

	// bowldOverImageConfig returns dimensions of bowldOverImagePath
	bowldOverImageConfig, _, err := image.DecodeConfig(util.GetReaderFromFile(bowldOverImagePath))
	if err != nil {
		catlog.Fatalf("Could not load bowldOver image %v", err)
	}

	// bowldOverImageDimensions is effectively the size of the image
	bowldOverImageDimensions = pixel.R(0, 0, float64(bowldOverImageConfig.Width), float64(bowldOverImageConfig.Height))

	// Set properties that require bounding box
	BowldOver.bounds = bowldOverImageDimensions

	// Load the background image
	bowldOverBackgroundSprite, bowldOverBackgroundPic = util.LoadSprite(bowldOverImagePath, bowldOverImageDimensions)

	// Get all collision bounds from the CSV file
	bowldOverCollisions = loadCollisions(bowldOverCollisionPath, BowldOver.bounds)
}

func initBowldOver() {
	player.SetPos(bowldOverStartPos)
}

func updateBowldOver(dt float64, win *pixelgl.Window) {
	// Try move the player
	if !movePlayer(win, dt, bowldOverCollisions) {
		// Player is not moving, update animation
		player.AnimateIdle()
	}
}

func drawBowldOver() {
	bowldOverBackgroundSprite.Draw(consts.GameView, pixel.IM.Moved(bowldOverImageDimensions.Center()))
}

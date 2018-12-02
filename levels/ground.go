package levels

import (
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

	backgroundSprite *pixel.Sprite
	backgroundPic    pixel.Picture

	// groundImageDimensions is effectively the size of the image
	groundImageDimensions = pixel.R(0, 0, 1000, 1000)

	// groundFloorCollisions are all the rectangles which should cause the player
	// to collide: i.e. unpassable
	groundFloorCollisions []pixel.Rect

	groundFloorStartPos = pixel.V(0, 0)
)

func init() {
	catlog.Debug("Preparing ground level")

	// Load the background image
	backgroundSprite, backgroundPic = util.LoadSprite(groundImagePath, groundImageDimensions)

	// Get all collision bounds from the CSV file
	groundFloorCollisions = loadCollisions(groundCollisionPath)
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
	// inverseMoved is the player position inversed
	backgroundSprite.Draw(target, pixel.IM)
}

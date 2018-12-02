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
	groundImagePath          = "assets/graphics/groundFloor.png"
	groundCollisionPath      = "assets/csv/groundFloorWalls.csv"
	groundActivationZonesCSV = "assets/csv/groundFloorZones.csv"
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

	groundZones     *map[pixel.Rect]string
	groundZoneFuncs = map[string]func(){}
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

	// Set properties that require bounding box
	Ground.bounds = groundImageDimensions
	groundZones = loadActivationZones(groundActivationZonesCSV, groundImageDimensions)

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

	// Check for activation zone changes
	zoneFunc := player.GetActivationZoneChange(*groundZones)
	if zoneFunc != "" {
		catlog.Debugf("Got new zone, trying to call function '%s'", zoneFunc)
		if f, ok := groundZoneFuncs[zoneFunc]; ok {
			f()
		} else {
			catlog.Debugf("Did not find function %s, doing nothing", zoneFunc)
		}
	}
}

func drawGround(target pixel.Target) {
	groundBackgroundSprite.Draw(target, pixel.IM.Moved(groundImageDimensions.Center()))
}

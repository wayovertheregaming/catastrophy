package levels

import (
	"image"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/consts"
	"github.com/wayovertheregaming/catastrophy/dialogue"
	"github.com/wayovertheregaming/catastrophy/player"
	"github.com/wayovertheregaming/catastrophy/riddles"
	"github.com/wayovertheregaming/catastrophy/trophies"
	"github.com/wayovertheregaming/catastrophy/util"
)

const (
	firstImagePath          = "assets/graphics/firstFloor.png"
	firstCollisionPath      = "assets/csv/firstFloorWalls.csv"
	firstActivationZonesCSV = "assets/csv/firstFloorZones.csv"
)

var (
	// First is the first floor level
	First = &Level{
		name:          consts.LevelNameFirst,
		updateFunc:    updateFirst,
		drawFunc:      drawFirst,
		initFunc:      initFirst,
		displayPlayer: true,
		musicFile:     "PlayingInHouse.mp3",
	}

	firstBackgroundSprite *pixel.Sprite
	firstBackgroundPic    pixel.Picture

	// firstFloorCollisions are all the rectangles which should cause the player
	// to collide: i.e. unpassable
	firstFloorCollisions []pixel.Rect
	firstImageDimensions pixel.Rect
	firstFloorStartPos   = pixel.V(1300, -1400)

	// firstZones holds the zones and the function name to call, as read from the
	// CSV
	firstZones *map[pixel.Rect]string
	// firstZoneFuncs is a map of function names (as they appear in the CSV) and
	// the function as defined in this file
	firstZoneFuncs = map[string]func(){
		"spider": speakToSpider,
	}

	spokenToSpider = false
)

func init() {
	catlog.Debug("Preparing first floor")

	//firstImageConfig returns dimensions of firstImagePath
	firstImageConfig, _, err := image.DecodeConfig(util.GetReaderFromFile(firstImagePath))
	if err != nil {
		catlog.Fatalf("Could not load first floor image %v", err)
	}

	// firstImageDimensions is effectively the size of the image
	firstImageDimensions = pixel.R(0, 0, float64(firstImageConfig.Width), float64(firstImageConfig.Height))

	// Set properties that require bounding box
	First.bounds = firstImageDimensions
	firstZones = loadActivationZones(firstActivationZonesCSV, firstImageDimensions)

	// Load the background image
	firstBackgroundSprite, firstBackgroundPic = util.LoadSprite(firstImagePath, firstImageDimensions)

	// Get all collision bounds from the CSV file
	firstFloorCollisions = loadCollisions(firstCollisionPath, First.bounds)
}

func initFirst() {
	player.SetPos(firstFloorStartPos)
}

func updateFirst(dt float64, win *pixelgl.Window) {
	// Try move the player
	if !movePlayer(win, dt, firstFloorCollisions) {
		// Player is not moving, update animation
		player.AnimateIdle()
	}

	// Check for activation zone changes
	zoneFunc := player.GetActivationZoneChange(*firstZones)
	if zoneFunc != "" {
		catlog.Debugf("Got new zone, trying to call function '%s'", zoneFunc)
		if f, ok := firstZoneFuncs[zoneFunc]; ok {
			f()
		} else {
			catlog.Debugf("Did not find function %s, doing nothing", zoneFunc)
		}
	}
}

func drawFirst() {
	firstBackgroundSprite.Draw(consts.GameView, pixel.IM.Moved(firstImageDimensions.Center()))
}

func speakToSpider() {
	if spokenToSpider {
		return
	}

	failDialogue := []dialogue.Dialogue{
		dialogue.Dialogue{
			IsPlayer: false,
			Name:     "Spider",
			Text:     "Sorry, wrong.\nTry again later",
		},
	}

	riddles.RunRiddle(
		dialogue.FirstSpiderRiddle,
		failDialogue,
		trophies.RidSpider,
		passedSpider,
	)
}

func passedSpider() {
	spokenToSpider = true
}

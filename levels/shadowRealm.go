package levels

import (
	"fmt"
	"image"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/consts"
	"github.com/wayovertheregaming/catastrophy/dialogue"
	"github.com/wayovertheregaming/catastrophy/player"
	"github.com/wayovertheregaming/catastrophy/trophies"
	"github.com/wayovertheregaming/catastrophy/util"
)

const (
	shadowImagePath          = "assets/graphics/shadowRealm.png"
	shadowActivationZonesCSV = "assets/csv/shadowZones.csv"
	catGodImagePath          = "assets/graphics/catGod.png"
)

var (
	// ShadowRealm is the ground level
	ShadowRealm = &Level{
		name:          consts.LevelNameShadow,
		updateFunc:    updateShadow,
		drawFunc:      drawShadow,
		initFunc:      initShadow,
		displayPlayer: true,
		musicFile:     "ShadowRealm.mp3",
	}

	shadowBackgroundSprite *pixel.Sprite

	// shadowImageDimensions is effectively the size of the image
	shadowImageDimensions pixel.Rect
	shadowFloorStartPos   = pixel.V(40, -1000)

	// shadowZones
	shadowZones     *map[pixel.Rect]string
	shadowZoneFuncs = map[string]func(){
		"talkToGod": talkToGod,
	}

	catGodSprite       *pixel.Sprite
	catGodShift        = pixel.ZV
	catGodShiftCounter float64

	firstVisit = true
)

func init() {
	catlog.Debug("Preparing shadow realm level")

	// Load the background image
	shadowBackgroundSprite, _ = util.LoadSprite(shadowImagePath, shadowImageDimensions)

	// shadowImageConfig returns dimensions of groundImagePath
	shadowImageConfig, _, err := image.DecodeConfig(util.GetReaderFromFile(groundImagePath))
	if err != nil {
		catlog.Fatalf("Could not load shadow realm image %v", err)
	}

	// shadowmageDimensions is effectively the size of the image
	shadowImageDimensions = pixel.R(0, 0, float64(shadowImageConfig.Width), float64(shadowImageConfig.Height))
	ShadowRealm.bounds = shadowImageDimensions
	shadowZones = loadActivationZones(shadowActivationZonesCSV, shadowImageDimensions)

	// Load the background image
	shadowBackgroundSprite, _ = util.LoadSprite(shadowImagePath, shadowImageDimensions)

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

	// Check for activation zone changes
	zoneFunc := player.GetActivationZoneChange(*shadowZones)
	if zoneFunc != "" {
		catlog.Debugf("Got new zone, trying to call function '%s'", zoneFunc)
		if f, ok := shadowZoneFuncs[zoneFunc]; ok {
			f()
		} else {
			catlog.Debugf("Did not find function %s, doing nothing", zoneFunc)
		}
	}
}

func drawShadow() {
	// inverseMoved is the player position inversed
	shadowBackgroundSprite.Draw(consts.GameView, pixel.IM.Moved(shadowImageDimensions.Center()))

	catGodSprite.Draw(
		consts.GameView,
		pixel.IM.Moved(shadowImageDimensions.Center().Add(catGodShift)).Scaled(shadowImageDimensions.Center(), 3),
	)
}

func talkToGod() {
	// Have to run this in go routine because it blocks
	// Each dialogue call to start waits for the channel to return
	go func() {
		if firstVisit {
			<-dialogue.Start(dialogue.FirstVisitToShadow)
			firstVisit = false
		}

		if len(player.GetInventory()) == 0 {
			<-dialogue.Start(dialogue.ShadowHaveNoItems)
		} else {
			<-dialogue.Start(dialogue.ShadowHaveItems)
			trophies.Sacrifice(player.GetInventory())
			// Remove all items - they've been sacrificed
			player.SacrificeAll()
		}

		<-dialogue.Start([]dialogue.Dialogue{
			{
				IsPlayer: false,
				Name:     "Cat god",
				Text:     fmt.Sprintf("There are %d things left to\ncollect", trophies.HowManyUnsacrificed),
			},
		})
		<-dialogue.Start(dialogue.ShadowExit)
	}()
}

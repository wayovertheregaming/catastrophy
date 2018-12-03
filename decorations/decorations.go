package decorations

import (
	"bytes"
	"encoding/csv"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/wayovertheregaming/catastrophy/assets"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/consts"
	"github.com/wayovertheregaming/catastrophy/gamestate"
	"github.com/wayovertheregaming/catastrophy/player"
	"github.com/wayovertheregaming/catastrophy/util"
)

const (
	decorationsImgPath = "assets/graphics/decorations.png"
	decorationsCSVPath = "assets/csv/decorations.csv"
)

var (
	// allSprites is all the sprites from the decorations spritemap. It is sub-
	// divided by level name
	allSprites = make(map[string]map[*pixel.Sprite]pixel.Vec)

	spritesheet pixel.Picture
)

func init() {
	catlog.Debug("Doing decorations init")

	// Load sprites
	spritesheet = util.LoadPic(decorationsImgPath)
	consts.DecorationsLayer = pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)

	loadSprites()
}

// Draw will draw all sprites to the screen based on what the level has set.
// This is done by checking the gamestate for the level name
func Draw() {
	currentLevel := gamestate.GetLevel().Name()

	// Do not attempt to draw if the level has no sprites
	if _, ok := allSprites[currentLevel]; !ok {
		return
	}

	// The offset for the decorations is the same as the map
	levelCentre := gamestate.GetLevel().Bounds().Center()
	// Offset is moving:
	//  - up-right to the centre of the window
	//  - down-left the distance the player has moved
	//  - down-left the centre of this level
	//  - up-right the centre of this sprite (this one done in for loop)
	cam := pixel.IM.Moved(consts.WinBounds.Center().Sub(player.GetPos()).Sub(levelCentre))

	// Loop through each sprite for this level
	for sprite, pos := range allSprites[currentLevel] {
		spriteCentre := sprite.Frame().Size().Scaled(0.5)
		sprite.Draw(consts.DecorationsLayer, cam.Moved(pos.Add(spriteCentre)))
	}
}

func loadSprites() {
	catlog.Debugf("Loading decorations CSV: %s", decorationsCSVPath)

	decorationsF, err := assets.Asset(decorationsCSVPath)
	if err != nil {
		catlog.Fatalf("Could not load CSV: %v", err)
	}

	csvReader := csv.NewReader(bytes.NewReader(decorationsF))
	decorations, err := csvReader.ReadAll()
	if err != nil {
		catlog.Fatalf("Could not read CSV: %v", err)
	}

	for _, row := range decorations {
		decorationName := row[0]
		levelName := row[1]

		// Check if a submap for the level exists
		if _, ok := allSprites[levelName]; !ok {
			catlog.Debugf("Creating map for sprites on %s", levelName)
			// this level not yet setup, create a map
			allSprites[levelName] = make(map[*pixel.Sprite]pixel.Vec)
		}

		// Get the bounds of the sprite in the spritemap
		spritePosx1 := mustParseFloat64(row[2])
		spritePosy1 := mustParseFloat64(row[3])
		spritePosx2 := mustParseFloat64(row[4])
		spritePosy2 := mustParseFloat64(row[5])
		spritePos := pixel.R(spritePosx1, spritePosy1, spritePosx2, spritePosy2)

		// Get the position to draw the decoration on the screen
		placementx := mustParseFloat64(row[6])
		placementy := mustParseFloat64(row[7])
		placement := pixel.V(placementx, placementy)

		catlog.Debugf("Loading decoration '%s'- %s, from %v on sheet, with position %v", decorationName, levelName, spritePos, placement)

		// Load sprite and add it to the map
		sprite := pixel.NewSprite(spritesheet, spritePos)
		allSprites[levelName][sprite] = placement
	}
}

// mustParseFloat64 uses `strconv.ParseFloat` and creates a fatal error if an
// error occurs
func mustParseFloat64(s string) float64 {
	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		catlog.Fatalf("Could not convert float64: %v", err)
	}

	return f64
}

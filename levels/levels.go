// Package levels holds all 'leveller' structs.  Each level and menu should be in
// it's own file, e.g. the level for the ground floor should be in the file
// `groundFloor.go` in this directory
package levels

import (
	"bytes"
	"encoding/csv"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wayovertheregaming/catastrophy/assets"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/gamestate"
	"github.com/wayovertheregaming/catastrophy/player"
)

// Level represents a playable level.  It implements the `gamestate.Leveller`
// interface
type Level struct {
	name       string
	updateFunc func(float64, *pixelgl.Window)
	drawFunc   func(pixel.Target)
	initFunc   func()
	// displayPlayer determines whether the level needs the player displayed
	displayPlayer bool
}

// Update will update the state of the level
func (l *Level) Update(dt float64, win *pixelgl.Window) {
	l.updateFunc(dt, win)

	// Update the player if displayed
	if l.displayPlayer {
		player.Update(dt)
	}
}

// Init will initialise the level.  The gamestate is provided so we can see
// where player is, etc
func (l *Level) Init() {
	catlog.Debugf("Initialising %s", l.Name())

	l.initFunc()
	gamestate.UnPauseGame()
}

// Draw will draw the level and contents to the target
func (l *Level) Draw(target pixel.Target) {
	l.drawFunc(target)
}

// Name will return the name of the level
func (l *Level) Name() string {
	return l.name
}

// Menu represents a clickable menu. It implements the `gamestate.Leveller`
// interface
type Menu struct {
	name       string
	updateFunc func(float64, *pixelgl.Window)
	drawFunc   func(pixel.Target)
	initFunc   func()
}

// Update will update the state of the menu
func (m *Menu) Update(dt float64, win *pixelgl.Window) {
	m.updateFunc(dt, win)
}

// Init will initialise the menu
func (m *Menu) Init() {
	catlog.Debugf("Initialising %s", m.Name())

	m.initFunc()
	gamestate.PauseGame()
}

// Draw will draw the menu to the target
func (m *Menu) Draw(target pixel.Target) {
	m.drawFunc(target)
}

// Name will return the name of the level
func (m *Menu) Name() string {
	return m.name
}

// loadCollisions will read each line of a CSV expecting four columns
// x1,y1,x2,y2; these are the bottom left and top right coordinates of the box
func loadCollisions(path string) []pixel.Rect {
	catlog.Debugf("Loading collision CSV: %s", path)

	// Get the CSV file from assets
	collisionF, err := assets.Asset(path)
	if err != nil {
		catlog.Fatalf("Could not load CSV: %v", err)
	}

	// Read it as a CSV, getting all rows
	csvReader := csv.NewReader(bytes.NewReader(collisionF))
	collisions, err := csvReader.ReadAll()
	if err != nil {
		catlog.Fatalf("Could not read CSV: %v", err)
	}

	// retRect is the slice to return
	retRect := make([]pixel.Rect, len(collisions))

	// Loop each row of the CSV
	for _, row := range collisions {
		// Get the coords of rect
		x1 := mustParseFloat64(row[0])
		y1 := mustParseFloat64(row[1])
		x2 := mustParseFloat64(row[2])
		y2 := mustParseFloat64(row[3])

		retRect = append(retRect, pixel.R(x1, y1, x2, y2))
	}

	return retRect
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

// movePlayer will attempt to move the player if the user is pressing the keys
// Returns if the player is moving - can be used to change animation
func movePlayer(win *pixelgl.Window, dt float64, collisions []pixel.Rect) bool {
	isMoving := false

	if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyUp) {
		player.WalkUp(dt, collisions)
		isMoving = true
	}
	if win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyDown) {
		player.WalkDown(dt, collisions)
		isMoving = true
	}
	if win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyLeft) {
		player.WalkLeft(dt, collisions)
		isMoving = true
	}
	if win.Pressed(pixelgl.KeyD) || win.Pressed(pixelgl.KeyRight) {
		player.WalkRight(dt, collisions)
		isMoving = true
	}

	return isMoving
}

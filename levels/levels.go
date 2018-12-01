// Package levels holds all 'leveller' structs.  Each level and menu should be in
// it's own file, e.g. the level for the ground floor should be in the file
// `groundFloor.go` in this directory
package levels

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/gamestate"
)

// Level represents a playable level.  It implements the `gamestate.Leveller`
// interface
type Level struct {
	name       string
	updateFunc func(float64, *pixelgl.Window)
	drawFunc   func(pixel.Target)
	initFunc   func()
}

// Update will update the state of the level
func (l *Level) Update(dt float64, win *pixelgl.Window) {
	l.updateFunc(dt, win)
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
}

// Update will update the state of the menu
func (m *Menu) Update(dt float64, win *pixelgl.Window) {
	m.updateFunc(dt, win)
}

// Init will initialise the menu
func (m *Menu) Init() {
	catlog.Debugf("Initialising %s", m.Name())

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

package gamestate

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wayovertheregaming/catastrophy/catlog"
)

// These are the type of run states a game can be in
const (
	RunStatePlaying int = iota
	RunStatePaused
)

var (
	gs gameState
)

// GameState is a singleton which holds global info for the game state
// This should not be access directly, but via the exported functions
type gameState struct {
	runState int
	level    Leveller
}

func init() {
	gs = gameState{
		runState: RunStatePaused,
		// TODO(need to set this to intro/main menu)
		level: nil,
	}
}

// IsPaused will return whether the game is in a paused state
func IsPaused() bool {
	return gs.runState == RunStatePaused
}

// IsPlaying will return whether the game is in a playing state
func IsPlaying() bool {
	return gs.runState == RunStatePlaying
}

// PauseGame will set the run state to paused.  This can be called regardless of
// the current state
func PauseGame() {
	catlog.Debug("Pausing game")
	gs.runState = RunStatePaused
}

// UnPauseGame will set the run state to playing.  This can be called regardless
// of the current state
func UnPauseGame() {
	catlog.Debug("Unpausing game")
	gs.runState = RunStatePlaying
}

// GetLevel will return the currently active level
func GetLevel() Leveller {
	return gs.level
}

// SetLevel will set the level provided to be the currently active level.  It
// will then call `Init()` on that level
func SetLevel(newLevel Leveller) {
	catlog.Debugf("Setting level to %s", newLevel.Name())
	gs.level = newLevel
	gs.level.Init()
}

// Update will update the currently active leveller
func Update(dt float64, win *pixelgl.Window) {
	gs.level.Update(dt, win)
}

// Draw will call draw on the currently active leveller
func Draw(target pixel.Target) {
	gs.level.Draw(target)
}

package player

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/wayovertheregaming/catastrophy/consts"
	"github.com/wayovertheregaming/catastrophy/gamestate"
	"github.com/wayovertheregaming/catastrophy/util"
)

const (
	// velocity is the speed at which the player moves
	// The number 10 is used as a placeholder, no idea how fast this is!
	velocity float64 = 500

	maxHunger       float64 = 100
	hungerRate      float64 = 2
	maxBladder      float64 = 100
	bladderFillRate float64 = 2
)

// These consts hold different animation states
const (
	animationStateWalking int = iota
	animationStateIdle
	animationStateSitting
	animationStateSleep
)

var (
	p player
)

// player represents the cat that the player controls
type player struct {
	animationState int
	// pos is the bottom left position of the player
	// Technically this controls the backgroundd
	pos pixel.Vec
	// direction is the angle the player is pointing
	// We'll only be dealing with right angles
	direction float64
	inventory []string
	hunger    float64
	bladder   float64
	// animationFrame is the index of the frame
	animationFrame int
	// animationCounter is the number of seconds since the animation frame changed
	animationCounter float64
	// currentZone is the activation zone the player currently resides in.  This
	// allows us to avoid calling an acitvation function more than once.  This is
	// a blank string when the player is not in a zone.
	currentZone string
}

// bounds returns the current bounding box of the player
// This is the pos plus the player size
func (p *player) bounds() pixel.Rect {
	return pixel.Rect{
		Min: p.pos,
		Max: p.pos.Add(consts.PlayerSize),
	}
}

// nextBounds returns the next bounding box of the player
// This is the next pos plus the player size
func (p *player) nextBounds(v pixel.Vec) pixel.Rect {
	return pixel.Rect{
		Min: v,
		Max: v.Add(consts.PlayerSize),
	}
}

// changeAnimationState will attempt to change the animation state.  This will
// reset the animation frame and counter if the state changes
func (p *player) changeAnimationState(newState int) {
	// If the current animation state is the same as the new requested state, just
	// return
	if p.animationState == newState {
		return
	}
	p.animationFrame = 0
	p.animationCounter = 0

	p.animationState = newState
}

// update will move animation forward
func (p *player) update(dt float64) {
	// Add the amount of seconds since last update
	p.animationCounter += dt

	// Check if we need to tick over frames
	if p.animationCounter > frameRate {
		p.animationCounter = 0
		p.animationFrame++
	}

	// If we've reached the end of the animation loop, reset frame to 0
	if p.animationFrame == len(stateToSprites(p.animationState)) {
		p.animationFrame = 0
	}
}

func (p *player) setZone(zoneName string) {
	p.currentZone = zoneName
}

func init() {
	p = player{
		animationState:   animationStateIdle,
		pos:              pixel.ZV,
		direction:        0,
		inventory:        []string{},
		hunger:           100,
		bladder:          0,
		animationFrame:   0,
		animationCounter: 0,
	}
}

// AnimateSleep will set the player to animate sleeping
func AnimateSleep() {
	p.changeAnimationState(animationStateSleep)
}

// AnimateSit will set the player to animate sitting
func AnimateSit() {
	p.changeAnimationState(animationStateSitting)
}

// AnimateIdle will set the player to animate idling/standing
func AnimateIdle() {
	p.changeAnimationState(animationStateIdle)
}

// AnimateWalk will set the player to animate walking
func AnimateWalk() {
	p.changeAnimationState(animationStateWalking)
}

// WalkUp will move the player upwards and animate them walking
func WalkUp(dt float64, collisionables []pixel.Rect) {
	AnimateWalk()
	p.direction = 0
	// nextPos is the potenial next position.  Use this to calculate if the player
	// will collide
	nextPos := p.pos.Add(pixel.V(0, dt*velocity))
	// Loop each collision to find out if we can move
	for _, r := range collisionables {
		if util.RectCollides(r, p.nextBounds(nextPos)) {
			// Player does collide with something
			// return without moving
			return
		}
	}

	// No collisions, move the player
	p.pos = nextPos
}

// WalkDown will move the player downwards and animate them walking
func WalkDown(dt float64, collisionables []pixel.Rect) {
	AnimateWalk()
	p.direction = math.Pi
	// nextPos is the potenial next position.  Use this to calculate if the player
	// will collide
	nextPos := p.pos.Sub(pixel.V(0, dt*velocity))
	// Loop each collision to find out if we can move
	for _, r := range collisionables {
		if util.RectCollides(r, p.nextBounds(nextPos)) {
			// Player does collide with something
			// return without moving
			return
		}
	}

	// No collisions, move the player
	p.pos = nextPos
}

// WalkLeft will move the player left and animate them walking
func WalkLeft(dt float64, collisionables []pixel.Rect) {
	AnimateWalk()
	p.direction = math.Pi / 2
	// nextPos is the potenial next position.  Use this to calculate if the player
	// will collide
	nextPos := p.pos.Sub(pixel.V(dt*velocity, 0))
	// Loop each collision to find out if we can move
	for _, r := range collisionables {
		if util.RectCollides(r, p.nextBounds(nextPos)) {
			// Player does collide with something
			// return without moving
			return
		}
	}

	// No collisions, move the player
	p.pos = nextPos
}

// WalkRight will move the player right and animate them walking
func WalkRight(dt float64, collisionables []pixel.Rect) {
	AnimateWalk()
	p.direction = (math.Pi * 3) / 2
	// nextPos is the potenial next position.  Use this to calculate if the player
	// will collide
	nextPos := p.pos.Add(pixel.V(dt*velocity, 0))
	// Loop each collision to find out if we can move
	for _, r := range collisionables {
		if util.RectCollides(r, p.nextBounds(nextPos)) {
			// Player does collide with something
			// return without moving
			return
		}
	}

	// No collisions, move the player
	p.pos = nextPos
}

// GetInventory will return the players current inventory
func GetInventory() []string {
	return p.inventory
}

// Draw draws the player to the target
func Draw(target pixel.Target) {
	spritepic := stateFrameToSprites(p.animationState, p.animationFrame)

	// Draw to the centre of the window
	// playerShift is how much to shift the player by so it sits in the middle of the window
	playerShift := gamestate.GetLevel().Bounds().Center().Add(p.pos)
	spritepic.sprite.Draw(target, pixel.IM.Moved(playerShift).Rotated(playerShift, p.direction))
}

// GetActivationZoneChange checks if the player is in a zone different to the
// last time it checked.  If the zone is the same, or it is not in a zone, this
// function will return a blank string, otherwise it returns the name of the
// zone
func GetActivationZoneChange(zones map[pixel.Rect]string) string {
	// Loop through each zone
	for r, f := range zones {
		// Check if we collide with the zone
		if util.RectCollides(r, p.bounds()) {
			// Check if we've already returned we're in that zone
			if f == p.currentZone {
				// Same zone as last check, do not activate again
				return ""
			}

			// New zone
			p.setZone(f)
			return f
		}
	}
	// No zone matched, return blank string
	p.setZone("")
	return ""
}

// GetPos returns the current player position.
// This is really the position of the background, as the player will always be
// in the centre, but conceptually it's clearer to have this as the player pos
func GetPos() pixel.Vec {
	return p.pos
}

// SetPos will set the players position.  This should be used when initiating a
// level so the player is at the start position
func SetPos(v pixel.Vec) {
	p.pos = v
}

// Update will update the player with things such as animiation frame
func Update(dt float64) {
	p.update(dt)
}

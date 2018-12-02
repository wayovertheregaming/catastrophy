package player

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/wayovertheregaming/catastrophy/consts"
	"github.com/wayovertheregaming/catastrophy/util"
)

const (
	// velocity is the speed at which the player moves
	// The number 10 is used as a placeholder, no idea how fast this is!
	velocity float64 = 100

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

	// playerSize is the width and height of the player, assuming the player is
	// facing upwards
	// Note, not tested, just placeholder sizes
	playerSize = pixel.V(10, 30)
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
}

// bounds returns the current bounding box of the player
// This is the pos plus the player size
func (p *player) bounds() pixel.Rect {
	return pixel.Rect{
		Min: p.pos,
		Max: p.pos.Add(playerSize),
	}
}

// nextBounds returns the next bounding box of the player
// This is the next pos plus the player size
func (p *player) nextBounds(v pixel.Vec) pixel.Rect {
	return pixel.Rect{
		Min: v,
		Max: v.Add(playerSize),
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
		p.animationFrame++
	}

	// If we've reached the end of the animation loop, reset frame to 0
	if p.animationFrame == len(stateToSprites(p.animationState)) {
		p.animationFrame = 0
	}
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
	p.direction = (math.Pi * 3) / 4
	p.pos.X -= dt * velocity
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
	p.direction = math.Pi / 2
	p.pos.X += dt * velocity
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

// CollidesWith determines whether the player collides with a rectangle
func CollidesWith(obj pixel.Rect) bool {
	// TODO(actually calculate this - must encoporate if we are facing left/right or up/down)
	return false
}

// GetInventory will return the players current inventory
func GetInventory() []string {
	return p.inventory
}

// Draw draws the player to the target
func Draw(target pixel.Target) {
	spritepic := stateFrameToSprites(p.animationState, p.animationFrame)
	// Draw to the centre of the window
	spritepic.sprite.Draw(target, pixel.IM.Moved(consts.WinCentre))
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

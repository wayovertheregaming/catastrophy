package player

import (
	"math"

	"github.com/faiface/pixel"
)

const (
	// velocity is the speed at which the player moves
	// The number 10 is used as a placeholder, no idea how fast this is!
	velocity float64 = 10

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
	pos            pixel.Vec
	// direction is the angle the player is pointing
	// We'll only be dealing with right angles
	direction float64
	inventory []string
	hunger    float64
	bladder   float64
}

func init() {
	p = player{
		animationState: animationStateIdle,
		pos:            pixel.ZV,
		direction:      0,
		inventory:      []string{},
		hunger:         100,
		bladder:        0,
	}
}

// AnimateSleep will set the player to animate sleeping
func AnimateSleep() {
	p.animationState = animationStateSleep
}

// AnimateSit will set the player to animate sitting
func AnimateSit() {
	p.animationState = animationStateSitting
}

// AnimateIdle will set the player to animate idling/standing
func AnimateIdle() {
	p.animationState = animationStateIdle
}

// AnimateWalk will set the player to animate walking
func AnimateWalk() {
	p.animationState = animationStateWalking
}

// WalkUp will move the player upwards and animate them walking
func WalkUp(dt float64) {
	AnimateWalk()
	p.direction = 0
	p.pos.Y += dt * velocity
}

// WalkDown will move the player downwards and animate them walking
func WalkDown(dt float64) {
	AnimateWalk()
	p.direction = math.Pi
	p.pos.Y -= dt * velocity
}

// WalkLeft will move the player left and animate them walking
func WalkLeft(dt float64) {
	AnimateWalk()
	p.direction = (math.Pi * 3) / 4
	p.pos.X -= dt * velocity
}

// WalkRight will move the player right and animate them walking
func WalkRight(dt float64) {
	AnimateWalk()
	p.direction = math.Pi / 2
	p.pos.X += dt * velocity
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

}

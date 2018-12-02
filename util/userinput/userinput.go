package userinput

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/wayovertheregaming/catastrophy/consts"
)

var (
	isActive bool
	typed    string

	userInp = make(chan string, 1)

	// backgroundBounds is the rect that represents the background that users type
	// into.  This is the full width of the window, but 1/3 of the height
	backgroundBounds = pixel.R(0, consts.WinHeight/3, consts.WinWidth, (consts.WinHeight*2)/3)
	backgroundColour = color.RGBA{0x8a, 0xc3, 0x6a, 0xdd}
)

// GetUserInput will wait for the user to press keys, displaying what they type
// in real time.  Function will return when the user presses enter.  If the user
// presses esc, the function will return a blank string.  This function blocks
// until a string is returned
func GetUserInput() string {
	isActive = true
	typed = ""

	userTyped := <-userInp
	return userTyped
}

// Update will detect user input and update the internal buffer
// This will set the return value when the user presses enter
func Update(win *pixelgl.Window) {
	if !isActive {
		return
	}

	// Return nothing if the user presses esc
	if win.JustPressed(pixelgl.KeyEscape) {
		isActive = false
		userInp <- ""
	}
	// Return the string typed in so far if the user presses enter
	if win.JustPressed(pixelgl.KeyEnter) {
		isActive = false
		userInp <- typed
	}

	// Add anything the user has typed to the end of the string
	typed += win.Typed()
}

// Draw will draw the input window to the screen and display what the user types
func Draw() {
	// Don't draw if not expecting input
	if !isActive {
		return
	}

	consts.ImdLayer.Color = backgroundColour
	consts.ImdLayer.Push(backgroundBounds.Min, backgroundBounds.Max)
	consts.ImdLayer.Rectangle(0)
}
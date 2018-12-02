package userinput

import "github.com/faiface/pixel/pixelgl"

var (
	isActive bool
	typed    string

	userInp = make(chan string, 1)
)

// GetUserInput will wait for the user to press keys, displaying what they type
// in real time.  Function will return when the user presses enter.  If the user
// presses esc, the function will return a blank string
func GetUserInput() string {
	isActive = true
	typed = ""

	userTyped := <-userInp
	return userTyped
}

// Update will detect user input and update the internal buffer
// This will set the return value when the user presses enter
func Update(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEscape) {
		isActive = false
		userInp <- ""
	}
	if win.JustPressed(pixelgl.KeyEnter) {
		isActive = false
		userInp <- typed
	}

	typed += win.Typed()
}

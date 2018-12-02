package dialogue

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/consts"
	"github.com/wayovertheregaming/catastrophy/gamestate"
	"golang.org/x/image/font/basicfont"
)

const (
	dialogueWidth   = consts.WinWidth / 3
	dialogueHeight  = consts.WinHeight / 5
	dialogueMargin  = 20
	dialoguePadding = 20
)

var (
	// leftPos is the position of the bottom-left of a left side panel
	// This is used for the non-player talker, or system notifications
	leftPos = pixel.V(dialogueMargin, dialogueMargin)
	// rightPos is the position of the bottom-left of a right side panel
	// This is used for the player talker
	rightPos = pixel.V(consts.WinWidth-dialogueWidth-dialogueMargin, dialogueMargin)
	// textShiftPos is the vector to shift text from the minPos
	textShiftPos = pixel.V(dialoguePadding, dialogueHeight-dialoguePadding)

	// currentDialogue is the dialogue currently playing
	currentDialogue []Dialogue
	// dialoguePanel is the index of dialogue currently playing
	dialoguePanel int

	dialoguePanelSize = pixel.V(dialogueWidth, dialogueHeight)

	textColour                            = color.RGBA{0x00, 0x00, 0xf1, 0x00}
	dialoguePanelBackgroundColour         = color.RGBA{0xf3, 0xff, 0xf1, 0x00}
	dialoguePanelBorderColour             = color.RGBA{0x00, 0x00, 0xf1, 0x00}
	borderThickness               float64 = 2

	// atlas contains the font to writing text to screen
	atlas *text.Atlas
)

func init() {
	// TODO(get a nice font)
	atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)

	// Set the panel to negative
	dialoguePanel = -1
}

// Dialogue is a single information either said to, or said by the player
// Sprite, Picture can be left as nil
type Dialogue struct {
	IsPlayer bool
	Sprite   *pixel.Sprite
	Picture  *pixel.Picture
	Name     string
	Text     string
}

func (d *Dialogue) draw(target pixel.Target) {
	// Draw the containing box
	var posMin pixel.Vec

	if d.IsPlayer {
		posMin = rightPos
	} else {
		posMin = leftPos
	}

	imd := imdraw.New(nil)
	// Background
	imd.Color = dialoguePanelBackgroundColour
	imd.Push(posMin, posMin.Add(dialoguePanelSize))
	imd.Rectangle(0)
	// Border
	imd.Color = dialoguePanelBorderColour
	imd.Push(
		posMin.Sub(pixel.V(borderThickness, borderThickness)),
		posMin.Add(dialoguePanelSize).Add(pixel.V(borderThickness, borderThickness)),
	)
	imd.Rectangle(borderThickness)
	imd.Draw(target)

	// Write text to screen
	// TODO(low priority - type this out letter by letter)
	text := text.New(posMin.Add(textShiftPos), atlas)
	text.Color = textColour
	fmt.Fprintf(text, "%s \n\n Press space to continue", d.Text)

	text.Draw(target, pixel.IM.Scaled(text.Orig, 2))
}

// Start will play a dialogue.  This will pause the game until the dialogue ends
func Start(dialogue []Dialogue) {
	catlog.Debug("Starting dialogue")
	gamestate.PauseGame()

	currentDialogue = dialogue
	dialoguePanel = 0
}

// Draw will draw the current panel
func Draw(target pixel.Target) {
	// We set the panel to a negative number when not displaying anything
	if dialoguePanel < 0 {
		return
	}

	// Do not continue if we cannot display the panel
	if dialoguePanel >= len(currentDialogue) {
		catlog.Debug("End of dialogue, exiting")

		gamestate.UnPauseGame()
		// Reset dialogue info
		dialoguePanel = -1
		currentDialogue = []Dialogue{}
		return
	}

	currentDialogue[dialoguePanel].draw(target)
}

// Update will display the current dialogue if one exists
// Taking a float64 so if we need to use dt in future not as many changes needed
// set to _ here to indicate we're not actually using it at the moment
func Update(_ float64, win *pixelgl.Window) {
	// We set the panel to a negative number when not displaying anything
	if dialoguePanel < 0 {
		return
	}

	// Check if the player has skipped to the next dialogue panel
	if win.JustPressed(pixelgl.KeySpace) {
		dialoguePanel++
		catlog.Debug("Moving onto next panel of dialogue")
	}
}

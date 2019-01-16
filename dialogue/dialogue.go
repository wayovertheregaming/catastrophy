package dialogue

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
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

	textColour                            = color.RGBA{0x00, 0x00, 0xf1, 0xff}
	dialoguePanelBackgroundColour         = color.RGBA{0xf3, 0xff, 0xf1, 0xff}
	dialoguePanelBorderColour             = color.RGBA{0x00, 0x00, 0xf1, 0xff}
	borderThickness               float64 = 2

	// atlas contains the font to writing text to screen
	atlas *text.Atlas

	hasStopped chan struct{}
)

func init() {
	catlog.Debug("Doing dialogue init")

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

func (d *Dialogue) draw() {
	// Draw the containing box
	var posMin pixel.Vec

	if d.IsPlayer {
		posMin = rightPos
	} else {
		posMin = leftPos
	}

	// Background
	consts.ImdLayer.Color = dialoguePanelBackgroundColour
	consts.ImdLayer.Push(posMin, posMin.Add(dialoguePanelSize))
	consts.ImdLayer.Rectangle(0)
	// Border
	consts.ImdLayer.Color = dialoguePanelBorderColour
	consts.ImdLayer.Push(
		posMin.Sub(pixel.V(borderThickness, borderThickness)),
		posMin.Add(dialoguePanelSize).Add(pixel.V(borderThickness, borderThickness)),
	)
	consts.ImdLayer.Rectangle(borderThickness)

	// Write textForScreen to screen
	// TODO(low priority - type this out letter by letter)
	textForScreen := text.New(posMin.Add(textShiftPos), atlas)
	textForScreen.Color = textColour

	// Include name if the dialogue has it set
	if d.Name == "" {
		// No name, just print textForScreen
		_, _ = fmt.Fprintf(textForScreen, "%s \n\n Press space to continue", d.Text)
	} else {
		_, _ = fmt.Fprintf(textForScreen, "%s:\n%s \n\n Press space to continue", d.Name, d.Text)
	}

	textForScreen.Draw(consts.TextLayer, pixel.IM.Scaled(textForScreen.Orig, 2))
}

// Start will play a dialogue.  This will pause the game until the dialogue ends
func Start(dialogue []Dialogue) chan struct{} {
	catlog.Debug("Starting dialogue")
	gamestate.PauseGame()

	// Create a channel to indicate when the dialogue completes
	hasStopped = make(chan struct{}, 1)

	currentDialogue = dialogue
	dialoguePanel = 0

	return hasStopped
}

// Draw will draw the current panel
func Draw() {
	// We set the panel to a negative number when not displaying anything
	if dialoguePanel < 0 {
		return
	}

	// Do not continue if we cannot display the panel
	if dialoguePanel >= len(currentDialogue) {
		catlog.Debug("End of dialogue, exiting")

		gamestate.UnPauseGame()
		close(hasStopped)
		// Reset dialogue info
		dialoguePanel = -1
		currentDialogue = []Dialogue{}
		return
	}

	currentDialogue[dialoguePanel].draw()
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

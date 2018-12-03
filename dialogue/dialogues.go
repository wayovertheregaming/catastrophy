package dialogue

// This file is for storing actual dialogues
// They should all be exported and no functions should be in this file

var (
	// GoingUpstairs tells the player they are going up the stairs
	GoingUpstairs = []Dialogue{
		Dialogue{
			IsPlayer: false,
			Text:     "You are going upstairs",
		},
	}
)

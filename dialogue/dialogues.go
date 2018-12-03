package dialogue

import (
	"fmt"

	"github.com/wayovertheregaming/catastrophy/trophies"
)

// This file is for storing actual dialogues
// They should all be exported and no functions should be in this file
// Each piece of text can be a maximum of 28 chars before needing an \n

var (
	// GoingUpstairs tells the player they are going up the stairs
	GoingUpstairs = []Dialogue{
		Dialogue{
			IsPlayer: false,
			Text:     "You are going upstairs",
		},
	}

	// FirstVisitToShadow is the dialogue which happens when the player first
	// visits the shadow realm
	FirstVisitToShadow = []Dialogue{
		Dialogue{
			IsPlayer: true,
			Text:     "Hello...?",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Unknown entity",
			Text:     "Hello little cat",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "Who are you? Where am I?",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "The cat god, and you're in\nthe shadow realm",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "Oh",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "I will accept sacrifices so\nyou can win the game",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "What game?",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "I mean...ummm...life...?",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "OKay then...",
		},
	}

	// ShadowHaveItems is the dialogue which happens when the player has items to
	// sacrifice
	ShadowHaveItems = []Dialogue{
		Dialogue{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "I see you have things for me",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "Don't I have choice in the\nmatter?",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "No, sacrifices must be made",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "Huh",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "I have accepted your items",
		},
	}

	// ShadowHaveNoItems is the dialogue which happens when the player has no
	// items to sacrifice
	ShadowHaveNoItems = []Dialogue{
		Dialogue{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "It looks like you have no\nitems for me",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "Complete challenges and\ngames to get items",
		},
	}

	// ShadowExit is the dialogue to play when leaving the shadow realm
	ShadowExit = []Dialogue{
		Dialogue{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     fmt.Sprintf("There are %d things left to\ncollect", trophies.HowManyUnsacrificed),
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "You can leave here the same\nway you came",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "By sleeping",
		},
	}
)

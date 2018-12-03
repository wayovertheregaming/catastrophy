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
			Text:     "The cat god, and you're in the shadow realm",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "Oh",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "I will accept sacrifices so you can win the game",
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
	ShadowHaveItems = []Dialogue{}

	// ShadowHaveNoItems is the dialogue which happens when the player has no
	// items to sacrifice
	ShadowHaveNoItems = []Dialogue{}
)

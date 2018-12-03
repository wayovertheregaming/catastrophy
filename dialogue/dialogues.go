package dialogue

// This file is for storing actual dialogues
// They should all be exported and no functions should be in this file
// Each piece of text can be a maximum of 28 chars before needing an \n

var (

	//HowToPlay is the opening dialogue explaining the game
	HowToPlay = []Dialogue{
		Dialogue{
			IsPlayer: false,
			Text:     "WASD to move, P to sleep\nPlay as a cat not as a Sheep\nHope u think the game's neat\nCheers, Rich & Ben",
		},
	}

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
			Text:     "You can leave here the same\nway you came",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "By sleeping",
		},
	}

	// FirstSpiderRiddle is the riddle dialogue the spider gives on the first floor
	FirstSpiderRiddle = []Dialogue{
		Dialogue{
			IsPlayer: false,
			Name:     "Spider",
			Text:     "Hi",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "Hello",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Spider",
			Text:     "I'm a spider",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "I know, it says on the panel",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Spider",
			Text:     "The what?",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "Nevermind, what do you want?",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Spider",
			Text:     "I have a riddle for you",
		},
	}

	// FirstWifeRiddle placeholder
	FirstWifeRiddle = []Dialogue{
		Dialogue{
			IsPlayer: false,
			Name:     "Owners wife",
			Text:     "Hello kitty\nAren't you a pwetty kitty",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Owners wife",
			Text:     "Yes you are!",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Owners wife",
			Text:     "How about you solve a\nriddle?",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "A riddle?  I'm a cat",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Owners wife",
			Text:     "I'm sorry, I can't understand\ncat",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Owners wife",
			Text:     "Answer the riddle...",
		},
	}

	// FirstSuitcaseRiddle placeholder
	FirstSuitcaseRiddle = []Dialogue{
		Dialogue{
			IsPlayer: false,
			Name:     "Suitcase",
			Text:     "Yo, I'm...",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "You're a suitcase",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Suitcase",
			Text:     "Yup.",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "Whatever",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Suitcase",
			Text:     "Riddle?",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "Sure",
		},
	}

	// FirstPCRiddle placeholder
	FirstPCRiddle = []Dialogue{
		Dialogue{
			IsPlayer: false,
			Name:     "Email",
			Text:     "To human,\nI think you're stupid",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "Typical human email",
		},
		Dialogue{
			IsPlayer: false,
			Name:     "Email",
			Text:     "Prove you're not by\nsolving this riddle",
		},
		Dialogue{
			IsPlayer: true,
			Text:     "So many riddles in this\nhouse",
		},
	}
)

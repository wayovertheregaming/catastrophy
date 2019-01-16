package dialogue

// This file is for storing actual dialogues
// They should all be exported and no functions should be in this file
// Each piece of text can be a maximum of 28 chars before needing an \n

var (

	// HowToPlay is the opening dialogue explaining the game
	HowToPlay = []Dialogue{
		{
			IsPlayer: false,
			Text:     "WASD to move, P to sleep\nPlay as a cat not as a Sheep\nHope u think the game's neat\nCheers, Rich & Ben",
		},
	}

	// GoingUpstairs tells the player they are going up the stairs
	GoingUpstairs = []Dialogue{
		{
			IsPlayer: false,
			Text:     "You are going upstairs",
		},
	}
	// GoingDownstairs tells the player they are going down the stairs
	GoingDownstairs = []Dialogue{
		{
			IsPlayer: false,
			Text:     "You are going downstairs",
		},
	}

	// FirstVisitToShadow is the dialogue which happens when the player first
	// visits the shadow realm
	FirstVisitToShadow = []Dialogue{
		{
			IsPlayer: true,
			Text:     "Hello...?",
		},
		{
			IsPlayer: false,
			Name:     "Unknown entity",
			Text:     "Hello little cat",
		},
		{
			IsPlayer: true,
			Text:     "Who are you? Where am I?",
		},
		{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "The cat god, and you're in\nthe shadow realm",
		},
		{
			IsPlayer: true,
			Text:     "Oh",
		},
		{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "I will accept sacrifices so\nyou can win the game",
		},
		{
			IsPlayer: true,
			Text:     "What game?",
		},
		{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "I mean...ummm...life...?",
		},
		{
			IsPlayer: true,
			Text:     "OKay then...",
		},
	}

	// ShadowHaveItems is the dialogue which happens when the player has items to
	// sacrifice
	ShadowHaveItems = []Dialogue{
		{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "I see you have things for me",
		},
		{
			IsPlayer: true,
			Text:     "Don't I have choice in the\nmatter?",
		},
		{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "No, sacrifices must be made",
		},
		{
			IsPlayer: true,
			Text:     "Huh",
		},
		{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "I have accepted your items",
		},
	}

	// ShadowHaveNoItems is the dialogue which happens when the player has no
	// items to sacrifice
	ShadowHaveNoItems = []Dialogue{
		{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "It looks like you have no\nitems for me",
		},
		{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "Complete challenges and\ngames to get items",
		},
	}

	// ShadowExit is the dialogue to play when leaving the shadow realm
	ShadowExit = []Dialogue{
		{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "You can leave here the same\nway you came",
		},
		{
			IsPlayer: false,
			Name:     "Cat god",
			Text:     "By sleeping",
		},
	}

	// FirstSpiderRiddle is the riddle dialogue the spider gives on the first floor
	FirstSpiderRiddle = []Dialogue{
		{
			IsPlayer: false,
			Name:     "Spider",
			Text:     "Hi",
		},
		{
			IsPlayer: true,
			Text:     "Hello",
		},
		{
			IsPlayer: false,
			Name:     "Spider",
			Text:     "I'm a spider",
		},
		{
			IsPlayer: true,
			Text:     "I know, it says on the panel",
		},
		{
			IsPlayer: false,
			Name:     "Spider",
			Text:     "The what?",
		},
		{
			IsPlayer: true,
			Text:     "Nevermind, what do you want?",
		},
		{
			IsPlayer: false,
			Name:     "Spider",
			Text:     "I have a riddle for you",
		},
	}

	// FirstWifeRiddle placeholder
	FirstWifeRiddle = []Dialogue{
		{
			IsPlayer: false,
			Name:     "Owners wife",
			Text:     "Hello kitty\nAren't you a pwetty kitty",
		},
		{
			IsPlayer: false,
			Name:     "Owners wife",
			Text:     "Yes you are!",
		},
		{
			IsPlayer: false,
			Name:     "Owners wife",
			Text:     "How about you solve a\nriddle?",
		},
		{
			IsPlayer: true,
			Text:     "A riddle?  I'm a cat",
		},
		{
			IsPlayer: false,
			Name:     "Owners wife",
			Text:     "I'm sorry, I can't understand\ncat",
		},
		{
			IsPlayer: false,
			Name:     "Owners wife",
			Text:     "Answer the riddle...",
		},
	}

	// FirstSuitcaseRiddle placeholder
	FirstSuitcaseRiddle = []Dialogue{
		{
			IsPlayer: false,
			Name:     "Suitcase",
			Text:     "Yo, I'm...",
		},
		{
			IsPlayer: true,
			Text:     "You're a suitcase",
		},
		{
			IsPlayer: false,
			Name:     "Suitcase",
			Text:     "Yup.",
		},
		{
			IsPlayer: true,
			Text:     "Whatever",
		},
		{
			IsPlayer: false,
			Name:     "Suitcase",
			Text:     "Riddle?",
		},
		{
			IsPlayer: true,
			Text:     "Sure",
		},
	}

	// FirstPCRiddle placeholder
	FirstPCRiddle = []Dialogue{
		{
			IsPlayer: false,
			Name:     "Email",
			Text:     "To human,\nI think you're stupid",
		},
		{
			IsPlayer: true,
			Text:     "Typical human email",
		},
		{
			IsPlayer: false,
			Name:     "Email",
			Text:     "Prove you're not by\nsolving this riddle",
		},
		{
			IsPlayer: true,
			Text:     "So many riddles in this\nhouse",
		},
	}

	// GroundTV1 is the dialogue shown first when walking in front of the TV
	GroundTV1 = []Dialogue{
		{
			IsPlayer: false,
			Name:     "Owner",
			Text:     "Move cat\nI can't see the tv",
		},
	}
	// GroundTV2 is the dialogue shown the second time when walking in front of the TV
	GroundTV2 = []Dialogue{
		{
			IsPlayer: false,
			Name:     "Owner",
			Text:     "Cat this is really annoying",
		},
	}
	// GroundTV3 is the dialogue shown the third time when walking in front of the TV
	GroundTV3 = []Dialogue{
		{
			IsPlayer: false,
			Name:     "Owner",
			Text:     "STOP DOING THAT",
		},
	}
)

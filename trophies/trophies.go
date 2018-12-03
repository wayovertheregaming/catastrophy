package trophies

var (
	// AllTrophies are all the tropies in the game
	AllTrophies = []*Trophy{
		&Trophy{Name: "Achievement: Toilet water"},
		&Trophy{Name: "Riddle: spider"},
		&Trophy{Name: "Achievement: Pillow talk"},
		&Trophy{Name: "Riddle: Wife"},
		&Trophy{Name: "Item: shoe"},
		&Trophy{Name: "Riddle: suitcase"},
		&Trophy{Name: "Phone"},
		&Trophy{Name: "Riddle: PC"},
		&Trophy{Name: "Item: tin of tuna"},
		&Trophy{Name: "Item: full bowl"},
		&Trophy{Name: "Achievement: Block TV"},
		&Trophy{Name: "Riddle: Magazine"},
		&Trophy{Name: "Item: Human food"},
		&Trophy{Name: "Riddle: Guest"},
		&Trophy{Name: "Achievement: littering"},
		&Trophy{Name: "Riddle: catapillar"},
	}
)

// Trophy could be an achievement or item the player gets from doing a challenge
type Trophy struct {
	Name       string
	Collected  bool
	Sacrificed bool
}

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
		&Trophy{Name: "Achievement: Block TV"},
		&Trophy{Name: "Riddle: Magazine"},
		&Trophy{Name: "Item: Human food"},
		&Trophy{Name: "Riddle: Guest"},
		&Trophy{Name: "Achievement: littering"},
		&Trophy{Name: "Riddle: catapillar"},
	}

	// HowManyUnsacrified is how many trohpies are left to be sacrificed
	HowManyUnsacrified = len(AllTrophies)
)

// Trophy could be an achievement or item the player gets from doing a challenge
type Trophy struct {
	Name       string
	Collected  bool
	Sacrificed bool
}

// Sacrifice will mark each trophy as sacrificed
func Sacrifice(trs []*Trophy) {
	sacrificedCount := 0

	// Use naive looping, only 20 items
	for _, t := range trs {
		for _, at := range AllTrophies {
			if at.Name == t.Name {
				at.Sacrificed = true
				sacrificedCount++
			}
		}
	}

	// Reduce count of items unsacrificed
	HowManyUnsacrified -= sacrificedCount
}

package trophies

// All trophies given in the game
// comment out any we don't have time to add in!!!
var (
	AchToilet   = &Trophy{Name: "Achievement: Toilet water"}
	RidSpider   = &Trophy{Name: "Riddle: spider"}
	AchPillow   = &Trophy{Name: "Achievement: Pillow talk"}
	RidWife     = &Trophy{Name: "Riddle: Wife"}
	ItemShoe    = &Trophy{Name: "Item: shoe"}
	RidSuitcase = &Trophy{Name: "Riddle: suitcase"}
	AchFired    = &Trophy{Name: "Achievement: You're fired"}
	RidPC       = &Trophy{Name: "Riddle: PC"}
	ItemTuna    = &Trophy{Name: "Item: tin of tuna"}
	AchTV       = &Trophy{Name: "Achievement: Block TV"}
	RidMag      = &Trophy{Name: "Riddle: Magazine"}
	ItemFood    = &Trophy{Name: "Item: Human food"}
	RidGuest    = &Trophy{Name: "Riddle: Guest"}
	AchLitter   = &Trophy{Name: "Achievement: littering"}
	RidCata     = &Trophy{Name: "Riddle: catapillar"}
)
var (
	// AllTrophies are all the tropies in the game
	AllTrophies = []*Trophy{}

	// HowManyUnsacrificed is how many trohpies are left to be sacrificed
	HowManyUnsacrificed = len(AllTrophies)
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
	HowManyUnsacrificed -= sacrificedCount
}

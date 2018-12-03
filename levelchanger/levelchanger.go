package levelchanger

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/wayovertheregaming/catastrophy/consts"
	"github.com/wayovertheregaming/catastrophy/gamestate"
	"github.com/wayovertheregaming/catastrophy/levels"
	"github.com/wayovertheregaming/catastrophy/player"
)

var (
	// beforeSleepLevel is where the player was before sleeping
	beforeSleepLevel gamestate.Leveller
	// beforeSleepPos is the position the player was in before sleeping
	beforeSleepPos pixel.Vec
)

// Sleep will transition the player between the shadow realm and back
func Sleep() {
	player.Sleep()

	// Run in go routine so updates continue (i.e. animations)
	go func() {
		time.Sleep(consts.SleepFor)

		// determine which way the player is going
		if gamestate.GetLevel().Name() == consts.LevelNameShadow {
			// Player is in shadow realm
			// Send back to where they were beforehand
			gamestate.SetLevel(beforeSleepLevel)
			player.SetPos(beforeSleepPos)
		} else {
			// Player is awake, send to shadow realm
			beforeSleepPos = player.GetPos()
			beforeSleepLevel = gamestate.GetLevel()

			gamestate.SetLevel(levels.ShadowRealm)
		}
	}()
}

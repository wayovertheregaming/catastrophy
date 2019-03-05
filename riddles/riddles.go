package riddles

import (
	"bytes"
	"encoding/csv"
	"strings"

	"github.com/wayovertheregaming/catastrophy/assets"
	"github.com/wayovertheregaming/catastrophy/catlog"
	"github.com/wayovertheregaming/catastrophy/dialogue"
	"github.com/wayovertheregaming/catastrophy/player"
	"github.com/wayovertheregaming/catastrophy/trophies"
	"github.com/wayovertheregaming/catastrophy/util/userinput"
)

const (
	path = "assets/csv/riddles.csv"
)

var (
	riddles []riddle
	counter = 0
)

type riddle struct {
	question  string
	answer    string
	beenAsked bool
}

func init() {
	catlog.Debug("Doing riddles init")

	catlog.Debugf("Loading riddles CSV: %s", path)

	// Get the CSV file from assets
	riddlesF, err := assets.Asset(path)
	if err != nil {
		catlog.Fatalf("Could not load CSV: %v", err)
	}

	// Read it as a CSV, getting all rows
	csvReader := csv.NewReader(bytes.NewReader(riddlesF))
	riddlesCSV, err := csvReader.ReadAll()
	if err != nil {
		catlog.Fatalf("Could not read CSV: %v", err)
	}

	// Loop each row of the CSV
	for _, row := range riddlesCSV {
		r := riddle{
			question: row[0],
			answer:   row[1],
		}

		riddles = append(riddles, r)
	}
}

// GetRiddle selects a riddle from list
func GetRiddle() (string, string) {
	if counter == len(riddles) {
		counter = 0
	}

	r := riddles[counter]
	counter++
	return r.question, r.answer
}

// RunRiddle will run a riddle and check the answer for you
func RunRiddle(initDialogue, failDialogue []dialogue.Dialogue, prize *trophies.Trophy, successFunc func()) {
	go func() {
		<-dialogue.Start(initDialogue)
		r, a := GetRiddle()

		dialogue.Start([]dialogue.Dialogue{
			{
				IsPlayer: false,
				Text:     r,
			},
		})

		userAns := strings.TrimSpace(userinput.GetUserInput())
		catlog.Debugf("User: %s, actual: %s", strings.ToLower(userAns), strings.ToLower(a))
		if strings.ToLower(userAns) == strings.ToLower(a) {
			// Answer is correct
			player.GiveItem(prize)
			successFunc()
			return
		}

		dialogue.Start(failDialogue)
	}()
}

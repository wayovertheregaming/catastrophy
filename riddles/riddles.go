package riddles

import (
	"bytes"
	"encoding/csv"

	"github.com/wayovertheregaming/catastrophy/assets"
	"github.com/wayovertheregaming/catastrophy/catlog"
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

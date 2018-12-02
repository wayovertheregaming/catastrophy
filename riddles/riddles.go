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
)

type riddle struct {
	question  string
	answer    string
	beenAsked bool
}

func init() {
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
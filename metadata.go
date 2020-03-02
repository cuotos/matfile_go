package matfile_go

import (
	"regexp"
	"strings"
	"time"
)

type Metadata struct {
	//Site       string
	//MatchID    string
	Player1    string
	Player1Elo string
	Player2    string
	Player2Elo string

	EventTime time.Time

	//Variation string
	//Unrated   bool
	//Crawford  bool
	//CubeLimit int
}

func parseMetadata(input []string, metadata *Metadata) error {

	var (
		year  int
		month time.Month
		day   int
		hour  int
		min   int
	)

	for _, l := range input {

		if strings.HasPrefix(l, `; [Player 1 "`) {
			metadata.Player1 = extractString(`; \[Player 1 "(.*)"\]`, l)
		}

		if strings.HasPrefix(l, `; [Player 2 "`) {
			metadata.Player2 = extractString(`; \[Player 2 "(.*)"\]`, l)
		}

		if strings.HasPrefix(l, `; [Player 1 Elo "`) {
			metadata.Player1Elo = extractString(`; \[Player 1 Elo "(.*)"\]`, l)
		}

		if strings.HasPrefix(l, `; [Player 2 Elo "`) {
			metadata.Player2Elo = extractString(`; \[Player 2 Elo "(.*)"\]`, l)
		}

		if strings.HasPrefix(l, `; [EventDate "`) {
			dateString := extractString(`; \[EventDate "(.*)"\]`, l)
			parsedDate, _ := time.Parse(`2006.01.02`, dateString)

			year = parsedDate.Year()
			month = parsedDate.Month()
			day = parsedDate.Day()

			hour = 1
			min = 1
		}

		if strings.HasPrefix(l, `; [EventTime`) {
			timeString := extractString(`; \[EventTime "(.*)"\]`, l)
			parsedTime, _ := time.Parse(`15.04`, timeString)

			hour = parsedTime.Hour()
			min = parsedTime.Minute()
		}
	}

	metadata.EventTime = time.Date(year, month, day, hour, min, 0, 0, time.UTC)

	return nil
}

func extractString(regex, input string) string {
	rgx := regexp.MustCompile(regex)
	return rgx.FindStringSubmatch(input)[1]
}

func getField(field string) string {
	return "test"
}

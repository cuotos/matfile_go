package matfile_go

import (
	"log"
	"regexp"
	"time"
)

type Metadata struct {
	Site       string
	MatchID    string
	Player1    string
	Player1Elo string
	Player2    string
	Player2Elo string

	EventTime time.Time

	Variation string
	Unrated   string
	Crawford  string
	CubeLimit string
}

func parseMetadataLines(input []string) (*Metadata, error) {

	var (
		year  int
		month time.Month
		day   int
		hour  int
		min   int
	)

	var metadata = &Metadata{}

	metaLineRegex := `^; \[(.*) "(.*)"\]$`
	reg := regexp.MustCompile(metaLineRegex)

	for _, l := range input {

		matches := reg.FindStringSubmatch(l)

		if matches != nil {

			switch f := matches[1:]; f[0] {
			case "Site":
				metadata.Site = f[1]
			case "Match ID":
				metadata.MatchID = f[1]
			case "Player 1":
				metadata.Player1 = f[1]
			case "Player 1 Elo":
				metadata.Player1Elo = f[1]
			case "Player 2":
				metadata.Player2 = f[1]
			case "Player 2 Elo":
				metadata.Player2Elo = f[1]
			case "EventDate":
				parsedDate, _ := time.Parse(`2006.01.02`, f[1])
				year = parsedDate.Year()
				month = parsedDate.Month()
				day = parsedDate.Day()
			case "EventTime":
				parsedTime, _ := time.Parse(`15.04`, f[1])
				hour = parsedTime.Hour()
				min = parsedTime.Minute()
			case "Variation":
				metadata.Variation = f[1]
			case "Unrated":
				metadata.Unrated = f[1]
			case "Crawford":
				metadata.Crawford = f[1]
			case "CubeLimit":
				metadata.CubeLimit = f[1]
			default:
				log.Printf("unknown metadata field, '%v: %v'", f[0], f[1])
			}
		} else {
			log.Printf(`no metadata found in line "%v"`, l)
		}
	}

	metadata.EventTime = time.Date(year, month, day, hour, min, 0, 0, time.UTC)

	return metadata, nil
}

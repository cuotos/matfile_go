package matfile_go

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Metadata struct {
	Player1 string
	Player1Elo string
	Player2 string
	Player2Elo string

	EventDate time.Time
	EventTime time.Time
}

func extractString(regex, input string) string {
	rgx := regexp.MustCompile(regex)

	return rgx.FindStringSubmatch(input)[1]
}

func parseMetadataString(line string, metadata *Metadata) error {
	if strings.HasPrefix(line, `; [Player 1 "`) {
		metadata.Player1 = extractString(`; \[Player 1 "(.*)"\]`, line)
	}

	if strings.HasPrefix(line, `; [Player 2 "`) {
		metadata.Player2 = extractString(`; \[Player 2 "(.*)"\]`, line)
	}

	if strings.HasPrefix(line, `; [Player 1 Elo `) {
		metadata.Player1Elo = extractString(`; \[Player 1 Elo "(.*)"\]`, line)
	}

	if strings.HasPrefix(line, `; [Player 2 Elo `) {
		metadata.Player2Elo = extractString(`; \[Player 2 Elo "(.*)"\]`, line)
	}

	if strings.HasPrefix(line, `; [EventDate `){
		eventDateString := extractString(`; \[EventDate "(.*)"\]`, line)
		fmt.Println(eventDateString)
		d, err := time.Parse("2006.01.02", eventDateString); if err != nil {
			return err
		} else {
			metadata.EventDate = d
		}
	}

	return nil
}
package matfile_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseSingleMetadataLines(t *testing.T) {
	tcs := []struct {
		inputLines []string
		expected   *Metadata
	}{
		{
			[]string{
				//`; [Site ""]`,
				//`; [Match ID "-618558043"]`,
				`; [Player 1 "Kubota Kanade"]`,
				`; [Player 2 "Moriuchi Toshiyuki"]`,
				`; [Player 1 Elo "1600.00/0"]`,
				`; [Player 2 Elo "1100.00/0"]`,
				`; [EventDate "2020.02.24"]`,
				`; [EventTime "13.54"]`,
				//`; [Variation "Backgammon"]`,
				//`; [Unrated "Off"]`,
				//`; [Crawford "On"]`,
				//`; [CubeLimit "1024"]`,
			},
			&Metadata{
				//Site:       "",
				//MatchID:    "-618558043",
				Player1:    "Kubota Kanade",
				Player2:    "Moriuchi Toshiyuki",
				Player1Elo: "1600.00/0",
				Player2Elo: "1100.00/0",
				EventTime:  time.Date(2020, 02, 24, 13, 54, 0, 0, time.UTC),
				//Variation:  "Backgammon",
				//Unrated:    false,
				//Crawford:   true,
				//CubeLimit:  1024,
			},
		},
	}

	for _, tc := range tcs {

		md := &Metadata{}

		err := parseMetadata(tc.inputLines, md)

		assert.NoError(t, err)
		assert.Equal(t, tc.expected, md)
	}
}

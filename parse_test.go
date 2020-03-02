package matfile_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseSingleMetadataLines(t *testing.T) {
	tcs := []struct {
		inputLine string
		expected *Metadata
	}{
		{
			`; [Player 1 "Kubota Kanade"]`,
			&Metadata{
				Player1: "Kubota Kanade",
			},
		},
		{
			`; [Player 2 "Moriuchi Toshiyuki"]`,
			&Metadata{
				Player2: "Moriuchi Toshiyuki",
			},
		},
		{
			`; [Player 1 Elo "1600.00/0"]`,
			&Metadata{
				Player1Elo: "1600.00/0",
			},
		},
		{
			`; [Player 2 Elo "1100.00/0"]`,
			&Metadata{
				Player2Elo: "1100.00/0",
			},
		},
		{
			`; [EventDate "2020.02.24"]`,
			&Metadata{
				EventDate: time.Date(2020,02,24,0,0,0,0,time.UTC),
			},
		},
	}

	for _, tc := range tcs {

		md := &Metadata{}

		err := parseMetadataString(tc.inputLine, md)

		assert.NoError(t, err)
		assert.Equal(t, tc.expected, md)
	}
}

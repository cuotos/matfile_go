package matfile_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	}

	for _, tc := range tcs {

		md := &Metadata{}

		parseMetadataString(tc.inputLine, md)

		assert.Equal(t, tc.expected, md)
	}
}

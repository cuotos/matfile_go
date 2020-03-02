package matfile_go

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestParseFullMatch(t *testing.T) {

	file, err := ioutil.ReadFile("testdata/fullmatch1.mat")
	if err != nil {
		t.Fatal(err)
	}

	match, err := ParseMatFile(string(file))

	assert.NoError(t, err)
	assert.Equal(t, 17, match.PointMatch)
}

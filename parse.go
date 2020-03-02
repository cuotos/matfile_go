package matfile_go

import (
	"github.com/spf13/cast"
	"strings"
)

type Match struct {
	Metadata   *Metadata
	PointMatch int
}

func ParseMatFile(file string) (*Match, error) {
	match := &Match{}

	metaLines := []string{}

	for _, line := range strings.Split(file, "\n") {
		if strings.HasPrefix(line, ";") {
			metaLines = append(metaLines, line)
		}

		if strings.HasSuffix(line, "point match") {
			pointMatchString := strings.TrimSuffix(line, " point match")
			match.PointMatch = cast.ToInt(pointMatchString)
		}
	}

	metadata, err := parseMetadataLines(metaLines)
	if err != nil {
		return nil, err
	}

	match.Metadata = metadata

	return match, nil
}

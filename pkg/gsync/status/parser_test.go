package status_test

import (
	"strings"
	"testing"

	"github.com/seblegall/gsync/pkg/gsync/status"
	"github.com/stretchr/testify/assert"
)

func TestParseShort(t *testing.T) {

	data := []struct {
		input string
		output status.Output
	}{
		{
			`
## master...origin/master
 M cmd/reset.go`,
			status.Output{
				Branch:      "master",
				FilesStatus: []string{
					"M cmd/reset.go",
				},
			},

		},
		{
			`## master...origin/master`,
			status.Output{
				Branch:      "master",
				FilesStatus: []string{},
			},

		},
	}


	for _, d := range data {
		assert.Equal(t, d.output.Branch, status.ParseShort(strings.NewReader(d.input)).Branch)
		assert.Equal(t, len(d.output.FilesStatus), len(status.ParseShort(strings.NewReader(d.input)).FilesStatus))
	}

}

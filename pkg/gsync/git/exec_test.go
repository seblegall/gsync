package git

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoteName(t *testing.T) {
	data := []struct{
		in string
		out string
	}{
		{
			"origin",
			"origin",
		},
		{
				`
origin
				`,
				"origin",
		},
		{
			`
origin
github
				`,
			"origin",
		},
	}

	for _, d := range data {
		execOutput = func(string) (io.Reader, error){ return strings.NewReader(d.in), nil }
		assert.Equal(t, d.out, RemoteName("test"))
	}
}


func TestDefaultBranch(t *testing.T) {

	RemoteName = func(string) string { return "origin"}

	data := []struct{
		in string
		out string
	}{
		{
			"origin/master",
			"master",
		},
		{
			`
origin/master
				`,
			"master",
		},
		{
			`
			
			
			
origin/master
			
			
			
				`,
			"master",
		},
		{
			`
origin/develop
				`,
			"develop",
		},
	}

	for _, d := range data {
		execOutput = func(string) (io.Reader, error){ return strings.NewReader(d.in), nil }
		assert.Equal(t, d.out, DefaultBranch("test"))
	}
}

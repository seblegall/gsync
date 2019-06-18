package schema_test

import (
	"github.com/Flaque/filet"
	"github.com/seblegall/gsync/pkg/gsync/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)


var (
	config = `
apiVersion: "gsync/v1alpha1"
project:
    name: "test"
    repositories:
        - git: "git@test.git"
          dir: "path/to/my/dir"
        - git: "git2@test.git"
          dir: "path/to/my/dir2"
`
	tmpfilename = "/tmp/gsync.yml"
)


func TestLoadProjects(t *testing.T) {
	defer filet.CleanUp(t)

	filet.File(t, tmpfilename, config)
	p, err := schema.LoadProject(tmpfilename)
	assert.Nil(t, err, "LoadProject failed with an error")
	assert.Equal(t, "test", p.Name)
	assert.Equal(t, "git@test.git", p.Repositories[0].Git)
	assert.Equal(t, "path/to/my/dir", p.Repositories[0].Dir)
}
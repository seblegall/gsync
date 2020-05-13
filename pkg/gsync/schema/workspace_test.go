package schema_test

import (
	"github.com/Flaque/filet"
	"github.com/seblegall/gsync/pkg/gsync/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)


var (
	tmpfilename = "/tmp/gsync.yml"
)

func TestLoadProject(t *testing.T) {

	config := `
apiVersion: "gsync/v1alpha1"
projects:
  - name: "getting-started"
    repositories:
      - git: "git@github.com:golang/example.git"
        dir: "./example"
      - git: "git@github.com:golang/blog.git"
        dir: "./blog"
`

	defer filet.CleanUp(t)

	filet.File(t, tmpfilename, config)
	projects, err := schema.LoadWorkspaces(tmpfilename)
	assert.Nil(t, err, "LoadWorkspaces failed with an error")
	assert.Equal(t, 1, len(projects))
	p := projects[0]
	assert.Equal(t, "getting-started", p.Name)
	assert.Equal(t, "git@github.com:golang/example.git", p.Repositories[0].Git)
	assert.Equal(t, "./example", p.Repositories[0].Dir)
}
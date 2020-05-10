package prompt

import (
	"github.com/seblegall/gsync/pkg/gsync/schema/v1alpha1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindReposNames(t *testing.T) {
	p := v1alpha1.Project{
		Name: "Test",
		Repositories: []v1alpha1.Repository{
			{
				"git@github.com/test/test1.git",
				"/path/to/my/dir",
				"",
				"",
			},
			{
				"git@github.com/test/test2.git",
				"/path/to/my/dir",
				"",
				"",
			},
			{
				"git@github.com/test/test3.git",
				"/path/to/my/dir",
				"",
				"",
			},
		},
	}
	names := findReposNames(p)

	assert.Equal(t, len(names), 3)
	assert.Equal(t, names[0], "test1")
	assert.Equal(t, names[1], "test2")
	assert.Equal(t, names[2], "test3")

}


func TestGetRepositoryFromNames(t *testing.T) {
	p := v1alpha1.Project{
		Name: "Test",
		Repositories: []v1alpha1.Repository{
			{
				"git@github.com/test/test1.git",
				"/path/to/my/dir",
				"",
				"",
			},
			{
				"git@github.com/test/test2.git",
				"/path/to/my/dir",
				"",
				"",
			},
			{
				"git@github.com/test/test3.git",
				"/path/to/my/dir",
				"",
				"",
			},
		},
	}

	repos := getRepositoryFromNames(p, []string{"test1"})

	assert.Equal(t, 1, len(repos))
	assert.Equal(t, v1alpha1.Repository{
		"git@github.com/test/test1.git",
		"/path/to/my/dir",
		"",
		"",
	}, repos[0])

}
package v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_GetName(t *testing.T) {
	repo := &Repository{
		"http://github.com/test/gitrepo.git",
		"/path/to/my/repo",
		"",
		"",
	}
	assert.Equal(t, "gitrepo", repo.Name())

	repo2 := &Repository{
		"git@github.com/test/gitrepo.git",
		"/path/to/my/repo",
		"",
		"",
	}
	assert.Equal(t,"gitrepo", repo2.Name())

	repo3 := &Repository{
		"git@github.com/test/gitrepo",
		"/path/to/my/repo",
		"",
		"",
	}
	assert.Equal(t, "", repo3.Name())


}
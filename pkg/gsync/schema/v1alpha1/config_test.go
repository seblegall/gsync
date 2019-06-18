package v1alpha1

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestRepository_GetName(t *testing.T) {
	repo := &Repository{
		"http://github.com/test/gitrepo.git",
		"/path/to/my/repo",
	}
	assert.Equal(t, repo.GetName(), "gitrepo")

	repo2 := &Repository{
		"git@github.com/test/gitrepo.git",
		"/path/to/my/repo",
	}
	assert.Equal(t, repo2.GetName(), "gitrepo")
}
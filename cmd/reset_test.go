package cmd

import (
	"testing"

	"github.com/seblegall/gsync/pkg/gsync/schema/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func TestValidateArgs(t *testing.T) {
	projects := []v1alpha1.Project{
		v1alpha1.Project{Name:"test"},
		v1alpha1.Project{Name:"test2"},
	}

	assert.Nil(t, validateArgs([]string{"test2"}, projects), "test2 could not be found in projects")
	assert.NotNil(t, validateArgs([]string{"test3"}, projects), "test3 should not be found in projects")
}
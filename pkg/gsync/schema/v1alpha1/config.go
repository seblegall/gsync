package v1alpha1

import (
	"github.com/seblegall/gsync/pkg/gsync/config"
	"regexp"
)

const (
	Version = "gsync/v1alpha1"
)

// NewGsyncConfig creates a GsyncConfig
func NewGsyncConfig() config.VersionedConfig {
	return new(GsyncConfig)
}

type GsyncConfig struct {
	APIVersion string      `yaml:"apiVersion"`
	Workspaces []Workspace `yaml:"workspaces"`
}

func (c *GsyncConfig) GetVersion() string {
	return c.APIVersion
}

type Workspace struct {
	Name string `yaml:"name"`
	Repositories []Repository `yaml:"repositories"`
}

type  Repository struct {
	Git string `yaml:"git"`
	Dir string `yaml:"dir"`
}

//Name returns the name of a git repository based on it's URL.
func (r *Repository) Name() string {
	re := regexp.MustCompile("/([^/]*)\\.git$")
	match := re.FindStringSubmatch(r.Git)
	if len(match) > 0 {
		return match[1]
	}

	return ""
}
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
	APIVersion string    `yaml:"apiVersion"`
	Projects []Workspace `yaml:"projects"`
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
	Remote string `yaml:"remote"`
	DefaultBranch string `yaml:"default-branch"`
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

//GetRemote returns the remote url ref for a given repository
//If no remote set, GetRemote will return "origin" as default remote url ref
func (r *Repository) GetRemote() string {
	if r.Remote == "" {
		return "origin"
	}

	return r.Remote
}

//GetDefaultBranch returns the default branch for a given repository
//If no default branch is set, it'll returns "master" as a default branch
func (r *Repository) GetDefaultBranch() string {
	if r.DefaultBranch == "" {
		return "master"
	}

	return r.DefaultBranch
}
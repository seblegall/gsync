package schema

import (
	"fmt"
	"github.com/seblegall/gsync/pkg/gsync/config"
	"github.com/seblegall/gsync/pkg/gsync/schema/v1alpha1"
	"gopkg.in/yaml.v2"
)

type APIVersion struct {
	Version string `yaml:"apiVersion"`
}

var SchemaVersions = Versions{
	{v1alpha1.Version, v1alpha1.NewGsyncConfig},
}

type Version struct {
	APIVersion string
	Factory    func() config.VersionedConfig
}

type Versions []Version

// Find search the constructor for a given api version.
func (v *Versions) Find(apiVersion string) (func() config.VersionedConfig, bool) {
	for _, version := range *v {
		if version.APIVersion == apiVersion {
			return version.Factory, true
		}
	}

	return nil, false
}

// ParseConfig reads a configuration file.
func ParseConfig(filename string) (config.VersionedConfig, error) {
	buf, err := config.ReadConfiguration(filename)
	if err != nil {
		return nil, err
	}

	apiVersion := &APIVersion{}
	if err := yaml.Unmarshal(buf, apiVersion); err != nil {
		return nil, err
	}

	factory, present := SchemaVersions.Find(apiVersion.Version)
	if !present {
		return nil, fmt.Errorf("unknown api version: '%s'", apiVersion.Version)
	}

	cfg := factory()
	if err := yaml.UnmarshalStrict(buf, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
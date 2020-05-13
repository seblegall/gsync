package schema

import (
	"github.com/mitchellh/go-homedir"
	"github.com/seblegall/gsync/pkg/gsync/schema/v1alpha1"
	"github.com/sirupsen/logrus"
)


func LoadWorkspaces(filename string) ([]v1alpha1.Workspace, error) {

	logrus.Debugf("loading project using %s", filename)

	parsed, err := ParseConfig(filename)
	if err != nil {
		return nil, err
	}

	config := parsed.(*v1alpha1.GsyncConfig)

	//Expand homedir
	for i, _ := range config.Workspaces {
		for y, _ := range config.Workspaces[i].Repositories {
			expDir, err := homedir.Expand(config.Workspaces[i].Repositories[y].Dir)
			if err != nil {
				continue
			}

			config.Workspaces[i].Repositories[y].Dir = expDir
		}
	}

	return config.Workspaces, nil
}
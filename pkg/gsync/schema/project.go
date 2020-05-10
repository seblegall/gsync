package schema

import (
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

	return config.Workspaces, nil
}
package schema

import (
	"github.com/seblegall/gsync/pkg/gsync/schema/v1alpha1"
	"github.com/sirupsen/logrus"
)


func LoadProject(filename string) (v1alpha1.Project, error) {

	logrus.Debugf("loading project using %s", filename)

	parsed, err := ParseConfig(filename)
	if err != nil {
		return v1alpha1.Project{}, err
	}

	config := parsed.(*v1alpha1.GsyncConfig)

	return config.Project, nil
}
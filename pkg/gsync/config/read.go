package config

import (
	"errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)


type VersionedConfig interface {
	GetVersion() string
}

func ReadConfiguration(filename string) ([]byte, error) {
	switch {
	case filename == "":
		return nil, errors.New("filename not specified")
	case filename == "-":
		return ioutil.ReadAll(os.Stdin)
	default:
		contents, err := ioutil.ReadFile(filename)
		if err != nil {
			if filename == "gsync.yaml" {
				logrus.Infof("Could not open gsync.yaml: \"%s\"", err)
				logrus.Infof("Trying to read from gsync.yml instead")
				contents, errIgnored := ioutil.ReadFile("gsync.yml")
				if errIgnored != nil {
					// Return original error because it's the one that matters
					return nil, err
				}

				return contents, nil
			}
		}

		return contents, err
	}
}

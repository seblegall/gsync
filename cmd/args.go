package cmd

import (
	"fmt"

	"github.com/seblegall/gsync/pkg/gsync/schema/v1alpha1"
)

func validateArgs(args []string, workspaces []v1alpha1.Workspace) error {

	for _, a := range args {
		found := false
		for _, p := range workspaces {
			if p.Name == a {
				found = true
			}
		}
		if found == false {
			return fmt.Errorf("arg %s doesn not match any referenced workspace", a)
		}
	}

	return nil
}
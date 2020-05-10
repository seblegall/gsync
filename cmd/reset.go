package cmd

import (
	"fmt"

	"github.com/seblegall/gsync/pkg/gsync/prompt"
	"github.com/seblegall/gsync/pkg/gsync/schema"
	"github.com/seblegall/gsync/pkg/gsync/schema/v1alpha1"
	"github.com/seblegall/gsync/pkg/gsync/sync"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// resetCmd represents the create command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "reset updates you project code with origin/master and discard not staged changes",
	Long: `the reset command clean up unstaged changes, checkout all repositories on master and rebase master on its origin branch`,
	Run: func(cmd *cobra.Command, args []string) {
		err := runReset(args)
		if err != nil {
			logrus.Fatal(err.Error())
		}
	},
}

func NewResetCommand() *cobra.Command {
	addInteractiveFlag(resetCmd)
	return resetCmd
}

func runReset(args []string) error {
	projects, err := schema.LoadProjects(configFile)
	if err != nil {
		return err
	}

	if err := validateArgs(args, projects); err != nil {
		return err
	}

	if len(projects) > 1 && len(args) == 0 {
		prompt.Warning("ℹ️ You have several projects configured. All of them will be reset.")
	}

	for _, p := range projects {
		if len(args) > 0 {
			for _, a := range args {
				if a == p.Name {
					if err := sync.Reset(p, interactive); err != nil {
						return err
					}
				}
			}
		} else {
			if err := sync.Reset(p, interactive); err != nil {
				return err
			}
		}

	}
	return nil
}

func validateArgs(args []string, projects []v1alpha1.Project) error {

	for _, a := range args {
		found := false
		for _, p := range projects {
			if p.Name == a {
				found = true
			}
		}
		if found == false {
			return fmt.Errorf("arg %s doesn not match any referenced project", a)
		}
	}

	return nil
}
package cmd

import (
	"github.com/seblegall/gsync/pkg/gsync/prompt"
	"github.com/seblegall/gsync/pkg/gsync/schema"
	"github.com/seblegall/gsync/pkg/gsync/sync"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init clone all the project repositories",
	Long: `the init command clone all the project repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		err := runInit(args)
		if err != nil {
			logrus.Fatal(err.Error())
		}
	},
}

func NewInitCommand() *cobra.Command {
	addInteractiveFlag(initCmd)
	return initCmd
}

func runInit(args []string) error {
	workspaces, err := schema.LoadWorkspaces(configFile)
	if err != nil {
		return err
	}

	if err := validateArgs(args, workspaces); err != nil {
		return err
	}

	if len(workspaces) > 1 && len(args) == 0 {
		prompt.Warning("ℹ️ You have several workspaces configured. All of them will be init.")
	}

	for _, w := range workspaces {
		if len(args) > 0 {
			for _, a := range args {
				if a == w.Name {
					if err := sync.Init(w, interactive); err != nil {
						return err
					}
				}
			}
		} else {
			if err := sync.Init(w, interactive); err != nil {
				return err
			}
		}

	}
	return nil
}
package cmd

import (
	"github.com/seblegall/gsync/pkg/gsync/schema"
	"github.com/seblegall/gsync/pkg/gsync/sync"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// statusCmd represents the create command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "status shows the status of each git repository from a workspace",
	Long: `the status command shows the status of each git repository from a workspace.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := runStatus(args)
		if err != nil {
			logrus.Fatal(err.Error())
		}
	},
}

func NewStatusCommand() *cobra.Command {
	addInteractiveFlag(statusCmd)
	return statusCmd
}

func runStatus(args []string) error {
	workspaces, err := schema.LoadWorkspaces(configFile)
	if err != nil {
		return err
	}

	if err := validateArgs(args, workspaces); err != nil {
		return err
	}

	for _, w := range workspaces {
		if len(args) > 0 {
			for _, a := range args {
				if a == w.Name {
					if err := sync.Status(w); err != nil {
						return err
					}
				}
			}
		} else {
			if err := sync.Status(w); err != nil {
				return err
			}
		}

	}
	return nil
}
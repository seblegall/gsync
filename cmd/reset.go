package cmd

import (
	"github.com/seblegall/gsync/pkg/gsync/prompt"
	"github.com/seblegall/gsync/pkg/gsync/schema"
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
	workspaces, err := schema.LoadWorkspaces(configFile)
	if err != nil {
		return err
	}

	if err := validateArgs(args, workspaces); err != nil {
		return err
	}

	if len(workspaces) > 1 && len(args) == 0 {
		prompt.Warning("ℹ️ You have several workspaces configured. All of them will be reset.")
	}

	for _, w := range workspaces {
		if len(args) > 0 {
			for _, a := range args {
				if a == w.Name {
					if err := sync.Reset(w, interactive); err != nil {
						return err
					}
				}
			}
		} else {
			if err := sync.Reset(w, interactive); err != nil {
				return err
			}
		}

	}
	return nil
}
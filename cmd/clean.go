package cmd

import (
	"github.com/seblegall/gsync/pkg/gsync/schema"
	"github.com/seblegall/gsync/pkg/gsync/sync"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cleanCmd represents the create command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean updates you project code with origin/master and discard not staged changes",
	Long: `the clean command clean up unstaged changes, checkout all repositories on master and rebase master on its origin branch`,
	Run: func(cmd *cobra.Command, args []string) {
		err := runClean()
		if err != nil {
			logrus.Fatal(err.Error())
		}
	},
}

func NewCleanCommand() *cobra.Command {
	return cleanCmd
}

func runClean() error {
	p, err := schema.LoadProject(filename)
	if err != nil {
		return err
	}

	return sync.Clean(p)
}
package sync

import (
	"fmt"

	"github.com/seblegall/gsync/pkg/gsync/git"
	"github.com/seblegall/gsync/pkg/gsync/prompt"
	"github.com/seblegall/gsync/pkg/gsync/schema/v1alpha1"
	"github.com/seblegall/gsync/pkg/gsync/status"
)

func Status(w v1alpha1.Workspace) error {
	prompt.Title(fmt.Sprintf("👀 status for workspace '%s'", w.Name))

	for _, r := range w.Repositories {

		output, err := git.ShortStatus(r.Dir)
		if err != nil {
			return err
		}

		defaultBranch := git.DefaultBranch(r.Dir)
		s := status.ParseShort(output)
		if len(s.FilesStatus) == 0 {
			if s.Branch == defaultBranch {
				prompt.InfoOK(fmt.Sprintf("✓ %s is on branch %s and clean", r.Name(), s.Branch))
			} else {
				prompt.Info(fmt.Sprintf("➤ %s is on branch %s and clean", r.Name(), s.Branch))
			}

		} else {
			prompt.InfoWarn(fmt.Sprintf("⚠︎ %s is on branch %s with uncommit changes", r.Name(), s.Branch))
			for _, fs := range s.FilesStatus {
				prompt.Info(fmt.Sprintf("    %s", fs))
			}
		}
	}

	return nil
}
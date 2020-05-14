package sync

import (
	"fmt"
	"os"

	"github.com/seblegall/gsync/pkg/gsync/git"
	"github.com/seblegall/gsync/pkg/gsync/prompt"
	"github.com/seblegall/gsync/pkg/gsync/schema/v1alpha1"
	log "github.com/sirupsen/logrus"
)

func setupRepositories(repos []v1alpha1.Repository) error {
	for _, repo := range repos {
		if _, err := os.Stat(repo.Dir); os.IsNotExist(err) {
			log.Infof("the directory '%s' does not exist", repo.Name())
			log.Infof("cloning '%s' into %s", repo.Name(), repo.Dir)
			if err := git.Clone(repo.Git, repo.Dir); err != nil {
				log.Errorf("git repository '%s' could not be cloned : %s", repo.Name(), err.Error() )
				log.Errorf("Unable to setup project entirely : %s", err.Error())
				return err
			}
		}
	}
	return nil
}

func Reset(w v1alpha1.Workspace, interactive bool) error {

	repos := w.Repositories

	prompt.Title(fmt.Sprintf("🙌 Resetting workspace %s", w.Name))

	if interactive {
		repos = prompt.SelectRepos(w)
	}

	if err := setupRepositories(repos); err != nil {
		log.Infof("Resetting only part of the project that has been successfully setup.")
	}

	for _, repo := range repos {

		defaultBranch := git.DefaultBranch(repo.Dir)
		remoteName := git.RemoteName(repo.Dir)

		if err := git.Reset(repo.Dir); err != nil {
			return err
		}
		if err := git.Fetch(repo.Dir); err != nil {
			return err
		}
		if err := git.Checkout(repo.Dir, defaultBranch); err != nil {
			return err
		}

		if err := git.Rebase(repo.Dir, fmt.Sprintf("%s/%s", remoteName, defaultBranch)); err != nil {
			return err
		}

		prompt.Info(fmt.Sprintf("'%s' is now on branch %s and up to date with %s", repo.Name(), defaultBranch, remoteName))

	}

	return nil
}


func Init(w v1alpha1.Workspace, interactive bool) error {

	repos := w.Repositories

	prompt.Title(fmt.Sprintf("🤙 Initing workspace %s", w.Name))

	if interactive {
		repos = prompt.SelectRepos(w)
	}

	if err := setupRepositories(repos); err != nil {
		log.Infof("Init only part of the workspace that has been successfully setup.")
	}

	for _, repo := range repos {

		defaultBranch := git.DefaultBranch(repo.Dir)
		remoteName := git.RemoteName(repo.Dir)

		prompt.Info(fmt.Sprintf("'%s' is now on branch %s and up to date with %s", repo.Name(), defaultBranch, remoteName))

	}

	return nil
}
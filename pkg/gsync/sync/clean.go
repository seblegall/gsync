package sync

import (
	"github.com/seblegall/gsync/pkg/gsync/git"
	"github.com/seblegall/gsync/pkg/gsync/schema/v1alpha1"
	"os"

	log "github.com/sirupsen/logrus"
)

func setupRepositories(p v1alpha1.Project) error {
	for _, repo := range p.Repositories {
		if _, err := os.Stat(repo.Dir); os.IsNotExist(err) {
			log.Infof("the directory '%s' does not exist", repo.GetName())
			log.Infof("cloning %s into %s", repo.GetName(), repo.Dir)
			if err := git.Clone(repo.Git, repo.Dir); err != nil {
				log.Errorf("git repository '%s' could not be cloned : %s", repo.GetName(), err.Error() )
				log.Errorf("Unable to setup project entirely : %s", err.Error())
				return err
			}
		}
	}
	return nil
}

func Clean(p v1alpha1.Project) error {

	log.Infof("cleaning up project %s", p.Name)

	if err := setupRepositories(p); err != nil {
		log.Infof("Cleaning up only part of the project that has been successfully setup.")
	}

	for _, repo := range p.Repositories {
		if err := git.Reset(repo.Dir); err != nil {
			return err
		}
		if err := git.Fetch(repo.Dir); err != nil {
			return err
		}
		if err := git.Checkout(repo.Dir, "master"); err != nil {
			return err
		}

		if err := git.Rebase(repo.Dir, "origin/master"); err != nil {
			return err
		}
	}

	log.Infof("project %s is now on branch master and up to date with origin", p.Name)


	return nil
}
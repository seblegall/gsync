package prompt

import (
	"github.com/seblegall/gsync/pkg/gsync/schema/v1alpha1"
	"github.com/AlecAivazis/survey/v2"
	"github.com/sirupsen/logrus"
)

func SelectRepos(p v1alpha1.Project) []v1alpha1.Repository {

	selected := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select repositories where to apply command :",
		Options: findReposNames(p),
	}
	err := survey.AskOne(prompt, &selected, survey.WithIcons(func(icons *survey.IconSet) {
		// you can set any icons
		icons.Question.Text = "✎"
		// for more information on formatting the icons, see here: https://github.com/mgutz/ansi#style-format
		icons.Question.Format = "cyan"
	}))

	if err != nil {
		if  err.Error() == "interrupt" {
			UserError("⛔️ Skipping project")
			return nil
		}

		logrus.Errorf("an error occurred when trying prompt : %s", err.Error())
		return nil
	}

	return getRepositoryFromNames(p, selected)
}

func getRepositoryFromNames(p v1alpha1.Project, names []string) []v1alpha1.Repository {
	var repos []v1alpha1.Repository
	for _, n := range names {
		for _, r := range p.Repositories {
			if r.Name() == n {
				repos = append(repos, r)
			}
		}
	}
	return  repos
}

func findReposNames(p v1alpha1.Project) []string {

	var names []string

	for _, r := range p.Repositories {
		names = append(names, r.Name())
	}

	return names
}

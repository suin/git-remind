package commands

import (
	"github.com/suin/git-remind/app/appservice"
	"github.com/suin/git-remind/app/cliutil"
	"github.com/urfave/cli"
)

var ReposCommand = cli.Command{
	Name:   "repos",
	Usage:  "Shows git repositories to be reminded",
	Action: reposAction,
}

func reposAction(*cli.Context) error {
	repos, err := appservice.GetRepos()
	if err != nil {
		return err
	}
	cliutil.PrintLines(repos)
	return nil
}

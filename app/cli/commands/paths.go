package commands

import (
	"github.com/suin/git-remind/app/appservice"
	"github.com/suin/git-remind/app/cliutil"
	"github.com/urfave/cli"
)

var PathsCommand = cli.Command{
	Name:   "paths",
	Usage:  "Shows path patterns configuration",
	Action: pathsAction,
}

func pathsAction(c *cli.Context) error {
	pathPatterns, err := appservice.GetPathPatterns()
	if err != nil {
		return err
	}
	cliutil.PrintLines(pathPatterns)
	return nil
}

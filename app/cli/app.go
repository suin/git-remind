package cli

import (
	"github.com/suin/git-remind/app/cli/cliglobalopts"
	"github.com/suin/git-remind/app/cli/commands"
	"github.com/urfave/cli"
)

var Name string
var Version string
var Description string

var App = &cli.App{
	Name:     Name,
	HelpName: Name,
	Usage:    Description,
	Version:  Version,
	Commands: cli.Commands{
		commands.PathsCommand,
		commands.ReposCommand,
		commands.StatusCommand,
		commands.StatusNotificationCommand,
	},
	Flags: []cli.Flag{
		cli.StringSliceFlag{
			Name:   "path",
			Usage:  "Path patterns of git repositories. If this options is specified, The configuration of `remind.paths` in .gitconfig is ignored. You may specify multiple paths like --path=/path/to/a --path=/path/to/b",
			EnvVar: "GIT_REMIND_PATHS",
		},
	},
	Before: func(context *cli.Context) error {
		cliglobalopts.SetPathPatterns(context.StringSlice("path"))
		return nil
	},
}

package cli

import (
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
}

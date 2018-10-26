package commands

import (
	"github.com/suin/git-remind/app/appservice"
	"github.com/urfave/cli"
)

var StatusNotificationCommand = cli.Command{
	Name:   "status-notification",
	Usage:  "Notifies repository statuses using desktop notification",
	Action: statusNotificationAction,
}

func statusNotificationAction(*cli.Context) error {
	err := appservice.NotifyRepoStatues()
	if err != nil {
		return err
	}
	return nil
}

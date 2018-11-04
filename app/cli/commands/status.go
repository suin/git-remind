package commands

import (
	"fmt"
	"github.com/suin/git-remind/app/appservice"
	"github.com/suin/git-remind/app/cliutil"
	"github.com/suin/git-remind/domain"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

var StatusCommand = cli.Command{
	Name:   "status",
	Usage:  "Shows repositories to need to git commit/push",
	Action: statusAction,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "short,s",
			Usage: "Give the output in the short-format",
		},
		cli.BoolFlag{
			Name:  "all,a",
			Usage: "Display all repositories status including up-to-date",
		},
		cli.BoolFlag{
			Name:  "no-status,n",
			Usage: "Display only repository paths",
		},
		cli.BoolFlag{
			Name:  "no-ansi",
			Usage: "Disable ANSI output",
		},
	},
}

func statusAction(c *cli.Context) error {
	repoStatuses, err := appservice.GetRepoStatuses()
	if err != nil {
		return err
	}
	color := getStatusColor(c)
	printStatus := getPrintStatus(c)
	for _, repoStatus := range repoStatuses {
		if !c.Bool("all") && repoStatus.GetGitStatus() == domain.UpToDate {
			continue
		}
		printStatus(repoStatus, color)
	}
	return nil
}

func getPrintStatus(c *cli.Context) func(domain.RepoStatus, cliutil.StatusColor) {
	if c.Bool("no-status") {
		return printOnlyPaths
	} else if c.Bool("short") {
		return printShortStatus
	}
	return printLongStatus
}

func getStatusColor(c *cli.Context) cliutil.StatusColor {
	if c.Bool("no-ansi") || !terminal.IsTerminal(int(os.Stdout.Fd())) {
		return &cliutil.StatusColorDisabled{}
	}
	return &cliutil.StatusColorEnabled{}
}

func printLongStatus(repoStatus domain.RepoStatus, color cliutil.StatusColor) {
	fmt.Printf("%s: %s\n", gitStatusToLongString(repoStatus.GetGitStatus(), color), repoStatus.GetPath())
}

func gitStatusToLongString(status domain.GitStatus, color cliutil.StatusColor) string {
	switch status {
	case domain.NeedToCommit:
		return color.NeedToCommit("Need to commit")
	case domain.NeedToPush:
		return color.NeedToPush("Need to push")
	case domain.NeedToCommitAndPush:
		return color.NeedToCommitAndPush("Need to commit and push")
	case domain.UpToDate:
		return color.UpToDate("Up-to-date")
	default:
		return color.Unknown("Unknown")
	}
}

func printShortStatus(repoStatus domain.RepoStatus, color cliutil.StatusColor) {
	fmt.Printf("%s %s\n", gitStatusToShortString(repoStatus.GetGitStatus(), color), repoStatus.GetPath())
}

func gitStatusToShortString(status domain.GitStatus, color cliutil.StatusColor) string {
	switch status {
	case domain.NeedToCommit:
		return color.NeedToCommit("C") + " "
	case domain.NeedToPush:
		return color.NeedToPush("P") + " "
	case domain.NeedToCommitAndPush:
		return color.NeedToCommit("C") + color.NeedToPush("P")
	case domain.UpToDate:
		return color.UpToDate("U") + " "
	default:
		return color.Unknown("??")
	}
}

func printOnlyPaths(repoStatus domain.RepoStatus, color cliutil.StatusColor) {
	fmt.Println(repoStatus.GetPath())
}

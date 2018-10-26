package infra

import (
	"github.com/suin/git-remind/domain"
	"github.com/suin/git-remind/infra/git"
	"regexp"
	"strings"
)

var GetGitStatus domain.GetGitStatus = func(repo string) (status domain.GitStatus, err error) {
	output, err := git.Status(repo)
	if err != nil {
		return
	}
	hasAheadCommits := hasAheadCommits(output)
	hasUncommittedFiles := hasUncommittedFiles(output)
	if hasUncommittedFiles && hasAheadCommits {
		return domain.NeedToCommitAndPush, err
	} else if hasAheadCommits {
		return domain.NeedToPush, err
	} else if hasUncommittedFiles {
		return domain.NeedToCommit, err
	} else {
		return domain.UpToDate, err
	}
}

func hasUncommittedFiles(output string) bool {
	return len(strings.Split(strings.Trim(output, "\r\n"), "\n")) > 1
}

func hasAheadCommits(output string) bool {
	ret, _ := regexp.MatchString(`^##.+\[ahead \d+]`, output)
	return ret
}

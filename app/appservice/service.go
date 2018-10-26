package appservice

import (
	"github.com/suin/git-remind/domain"
	"github.com/suin/git-remind/infra"
)

var GetPathPatterns = domain.NewGetPathPatterns(infra.GitGlobalConfigGetPathPatterns)
var GetRepos = domain.NewGetRepos(GetPathPatterns, infra.FilesystemRepos)
var GetRepoStatuses = domain.NewGetRepoStatuses(GetRepos, infra.GetGitStatus)
var NotifyRepoStatues = domain.NewNotifyRepoStatuses(GetRepoStatuses, infra.BeeepRepoStatusNotifier)

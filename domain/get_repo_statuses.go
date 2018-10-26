package domain

type GitStatus string

const (
	Unknown             GitStatus = "Unknown"
	NeedToCommit                  = "NeedToCommit"
	NeedToPush                    = "NeedToPush"
	NeedToCommitAndPush           = "NeedToCommitAndPush"
	UpToDate                      = "UpToDate"
)

type GetGitStatus func(repo string) (gitStatus GitStatus, err error)

type RepoStatus interface {
	GetPath() string
	GetGitStatus() GitStatus
}

type GetRepoStatuses func() (repoStatuses []RepoStatus, err error)

func NewGetRepoStatuses(getRepos GetRepos, getGitStatus GetGitStatus) GetRepoStatuses {
	return func() (repoStatuses []RepoStatus, err error) {
		repos, err := getRepos()
		for _, repo := range repos {
			gitStatus, err := getGitStatus(repo)
			if err != nil {
				continue
			}
			repoStatuses = append(repoStatuses, newRepoStatus(repo, gitStatus))
		}
		return
	}
}

type repoStatus struct {
	path   string
	status GitStatus
}

func (s *repoStatus) GetPath() string {
	return s.path
}

func (s *repoStatus) GetGitStatus() GitStatus {
	return s.status
}

func newRepoStatus(path string, status GitStatus) *repoStatus {
	return &repoStatus{path, status}
}

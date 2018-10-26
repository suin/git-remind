package domain

type RepoStatusNotifier interface {
	NotifyNeedToCommit(path string) (err error)
	NotifyNeedToPush(path string) (err error)
	NotifyNeedToCommitAndPush(path string) (err error)
}

type NotifyRepoStatuses func() (err error)

func NewNotifyRepoStatuses(getStatuses GetRepoStatuses, notifier RepoStatusNotifier) NotifyRepoStatuses {
	return func() (err error) {
		repoStatuses, err := getStatuses()
		if err != nil {
			return
		}
		for _, repoStatus := range repoStatuses {
			getNotificationMethod(notifier, repoStatus.GetGitStatus())(repoStatus.GetPath())
		}
		return
	}
}

func getNotificationMethod(notifier RepoStatusNotifier, status GitStatus) func(path string) error {
	switch status {
	case NeedToCommit:
		return notifier.NotifyNeedToCommit
	case NeedToPush:
		return notifier.NotifyNeedToPush
	case NeedToCommitAndPush:
		return notifier.NotifyNeedToCommitAndPush
	default:
		return func(path string) error {
			return nil
		}
	}
}

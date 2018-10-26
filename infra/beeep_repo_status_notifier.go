package infra

import (
	"github.com/gen2brain/beeep"
)

var BeeepRepoStatusNotifier = &beeepRepoStatusNotifier{}

type beeepRepoStatusNotifier struct {
}

func (*beeepRepoStatusNotifier) NotifyNeedToCommit(path string) (err error) {
	return beeep.Notify("Remind to git commit", path, "")
}

func (*beeepRepoStatusNotifier) NotifyNeedToPush(path string) (err error) {
	return beeep.Notify("Remind to git push", path, "")
}

func (*beeepRepoStatusNotifier) NotifyNeedToCommitAndPush(path string) (err error) {
	return beeep.Notify("Remind to git commit/push", path, "")
}

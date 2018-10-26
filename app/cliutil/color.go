package cliutil

import "github.com/mgutz/ansi"

type StatusColor interface {
	NeedToCommit(string) string
	NeedToPush(string) string
	NeedToCommitAndPush(string) string
	UpToDate(string) string
	Unknown(string) string
}

type StatusColorEnabled struct {
}

func (StatusColorEnabled) NeedToCommit(string string) string {
	return ansi.Color(string, "yellow")
}

func (StatusColorEnabled) NeedToPush(string string) string {
	return ansi.Color(string, "red")
}

func (StatusColorEnabled) NeedToCommitAndPush(string string) string {
	return ansi.Color(string, "red")
}

func (StatusColorEnabled) UpToDate(string string) string {
	return ansi.Color(string, "green")
}

func (StatusColorEnabled) Unknown(string string) string {
	return ansi.Color(string, "cyan")
}

type StatusColorDisabled struct {
}

func (StatusColorDisabled) NeedToCommit(string string) string {
	return string
}

func (StatusColorDisabled) NeedToPush(string string) string {
	return string
}

func (StatusColorDisabled) NeedToCommitAndPush(string string) string {
	return string
}

func (StatusColorDisabled) UpToDate(string string) string {
	return string
}

func (StatusColorDisabled) Unknown(string string) string {
	return string
}

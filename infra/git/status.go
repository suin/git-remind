package git

import (
	"os"
	"os/exec"
)

func Status(path string) (output string, err error) {
	preservedCwd, err := os.Getwd()
	if err != nil {
		return
	}
	defer os.Chdir(preservedCwd)
	err = os.Chdir(path)
	if err != nil {
		return
	}
	commandOutput, err := exec.Command("git", "status", "-sb").Output()
	if err != nil {
		return
	}
	output = string(commandOutput)
	return
}

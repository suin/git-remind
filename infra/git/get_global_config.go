package git

import (
	"os/exec"
	"strings"
)

func GetGlobalConfig(key string) (value string, err error) {
	output, err := exec.Command("git", "config", "--global", key).Output()
	if err != nil {
		return
	}
	value = strings.Trim(string(output), "\r\n")
	return
}

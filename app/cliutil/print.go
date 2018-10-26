package cliutil

import (
	"fmt"
	"strings"
)

func PrintLines(lines []string) {
	fmt.Println(strings.Join(lines, "\n"))
}

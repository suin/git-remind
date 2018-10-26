package main

import (
	"github.com/suin/git-remind/app/cli"
	"os"
)

func main() {
	cli.App.Run(os.Args)
}

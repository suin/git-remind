package infra

import (
	"github.com/suin/git-remind/domain"
	"os"
	"path/filepath"
)

var FilesystemRepos domain.GetReposByPathPattern = func(patterns domain.GetPathPatterns) (repos []string, err error) {
	pathPatterns, err := patterns()
	if err != nil {
		return
	}
	for _, pathPattern := range pathPatterns {
		paths, err2 := filepath.Glob(string(pathPattern))
		if err2 != nil {
			return repos, err2
		}
		for _, path := range paths {
			isDir, _ := isDirectory(path + "/.git")
			if isDir {
				repos = append(repos, path)
			}
		}
	}
	return
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

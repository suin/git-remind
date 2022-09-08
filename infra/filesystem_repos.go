package infra

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/suin/git-remind/domain"
)

var FilesystemRepos domain.GetReposByPathPattern = func(patterns domain.GetPathPatterns) (repos []string, err error) {
	pathPatterns, err := patterns()
	if err != nil {
		return
	}
	for _, pathPattern := range pathPatterns {
		if strings.HasSuffix(string(pathPattern), "/**") {
			path := strings.TrimSuffix(string(pathPattern), "/**")
			dirs, err := searchForGitDirs(path)
			if err != nil {
				return repos, err
			}
			repos = append(repos, dirs...)
		}
		paths, err := filepath.Glob(string(pathPattern))
		if err != nil {
			return repos, err
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

func searchForGitDirs(root string) ([]string, error) {
	var skipPath string
	var matchedPaths []string
	err := filepath.WalkDir(
		root,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() && d.Name() == ".git" && skipPath != path {
				skipPath = path
				matchedPaths = append(matchedPaths, strings.TrimSuffix(path, "/.git"))
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	return matchedPaths, nil
}

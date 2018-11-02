package cliglobalopts

import "github.com/suin/git-remind/domain"

var pathPatterns []string

func SetPathPatterns(p []string) {
	pathPatterns = p
}

var GetPathPatterns domain.GetPathPatterns = func() ([]string, error) {
	return pathPatterns, nil
}

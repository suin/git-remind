package domain

type GetReposByPathPattern func(patterns GetPathPatterns) (repos []string, err error)

type GetRepos func() (repos []string, err error)

func NewGetRepos(patterns GetPathPatterns, getRepos GetReposByPathPattern) GetRepos {
	return func() ([]string, error) {
		return getRepos(patterns)
	}
}

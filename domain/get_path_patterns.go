package domain

type GetPathPatterns func() (pathPatterns []string, err error)

func NewGetPathPatterns(pathPatterns GetPathPatterns) GetPathPatterns {
	return func() ([]string, error) {
		return pathPatterns()
	}
}

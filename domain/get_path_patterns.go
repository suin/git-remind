package domain

type GetPathPatterns func() (pathPatterns []string, err error)

func NewGetPathPatterns(pathPatterns GetPathPatterns) GetPathPatterns {
	return func() ([]string, error) {
		return pathPatterns()
	}
}

func MultipleGetPathPatterns(multipleGetPathPatterns ...GetPathPatterns) GetPathPatterns {
	return func() ([]string, error) {
		for _, getPathPatterns := range multipleGetPathPatterns {
			pathPatterns, err := getPathPatterns()
			if err != nil {
				return []string{}, err
			}
			if len(pathPatterns) > 0 {
				return pathPatterns, nil
			}
		}
		return []string{}, nil
	}
}

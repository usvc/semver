package semver

type Semvers []Semver

func (s Semvers) Len() int {
	return len(s)
}

func (s Semvers) Less(i, j int) bool {
	return !s[i].IsGreaterThan(s[j])
}

func (s Semvers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

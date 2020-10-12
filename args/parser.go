package args

import (
	"regexp"
)

const empty string = ""

func isArgument(n string) (m bool) {
	m, _ = regexp.MatchString("--?[a-z]+", n)
	return m
}

func getValueByIndex(i int, a []string) string {
	next := i + 1
	if next >= len(a) {
		return empty
	}
	v := a[next]
	if isArgument(v) {
		return empty
	}
	return v
}

func getAllValues(n string, a []string) (v []string) {
	for i, name := range a {
		if name == n {
			v = append(v, getValueByIndex(i, a))
		}
	}
	return v
}

// ParseArgs ...
func ParseArgs(a []string, n ...string) (p map[string][]string) {
	p = map[string][]string{}
	for _, name := range a {
		p[name] = getAllValues(name, a)
	}
	return p
}

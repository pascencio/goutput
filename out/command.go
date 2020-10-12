package out

import (
	"fmt"
	"os"
	"strings"
)

// Command ...
type Command struct {
	Level        Level
	Message      string
	Placeholders []string
}

func formatWithPlaceholders(m string, p []string) (f string) {
	for _, ph := range p {
		f = strings.Replace(m, "{}", ph, 1)
	}
	return f
}

// Print ...
func Print(c Command) (m string, s *os.File) {
	f := formatWithPlaceholders(c.Message, c.Placeholders)
	m = fmt.Sprintf("[%s]: %s\n", c.Level.GetName(), f)
	s = os.Stdout
	error := LevelError{}
	if c.Level.GetValue() == error.GetValue() {
		s = os.Stderr
	}
	fmt.Fprintf(s, m)
	return m, s
}

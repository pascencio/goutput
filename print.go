package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pascencio/goutput/args"
	"github.com/pascencio/goutput/out"
)

func formatWithPlaceholders(m string, p []string) (f string) {
	f = strings.Trim(m, "\"")
	for _, ph := range p {
		f = strings.Replace(f, "{}", strings.Trim(ph, "\""), 1)
	}
	return f
}

// PrintWithArgs ...
func PrintWithArgs(a []string) (m string, s *os.File) {
	p := args.ParseArgs(a, "-d", "-i", "-e", "-w", "-m", "-p")
	l, _ := out.ArgumentToLevel(p)
	c := out.Command{
		Level:        l,
		Message:      p["-m"][0],
		Placeholders: p["-p"],
	}
	return Print(c)
}

// Print ...
func Print(c out.Command) (m string, s *os.File) {
	f := formatWithPlaceholders(c.Message, c.Placeholders)
	m = fmt.Sprintf("[%s]: %s\n", c.Level.GetName(), f)
	s = os.Stdout
	error := out.LevelError{}
	if c.Level.GetValue() == error.GetValue() {
		s = os.Stderr
	}
	fmt.Fprintf(s, m)
	return m, s
}

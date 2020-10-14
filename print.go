package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pascencio/goutput/args"
	"github.com/pascencio/goutput/out"
)

// VERSION ...
const VERSION = "no-version"

func formatWithPlaceholders(m string, p []string) (f string) {
	f = strings.Trim(m, "\"")
	for _, ph := range p {
		f = strings.Replace(f, "{}", strings.Trim(ph, "\""), 1)
	}
	return f
}

// PrintVersion ...
func PrintVersion() (m string, s *os.File) {
	m = fmt.Sprintf("Goutput version %s", VERSION)
	s = os.Stdout
	fmt.Fprintf(s, m)
	return m, s
}

// PrintWithArgs ...
func PrintWithArgs(a []string) (m string, s *os.File) {
	p := args.ParseArgs(a, "-d", "-i", "-e", "-w", "-m", "-p", "--version")
	if _, ok := p["--version"]; ok && len(p) == 1 {
		return PrintVersion()
	}
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
	if !out.MustBePrinted(c.Level) {
		return "", nil
	}
	f := formatWithPlaceholders(c.Message, c.Placeholders)
	d := time.Now().Format("2006-01-01 15:04:05.000")
	m = fmt.Sprintf("[%s] - [%-5s]: %s\n", d, c.Level.GetName(), f)
	s = os.Stdout
	error := out.LevelError{}
	if c.Level.GetValue() == error.GetValue() {
		s = os.Stderr
	}
	fmt.Fprintf(s, m)
	return m, s
}

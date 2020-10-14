package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/pascencio/goutput/out"
)

const outformat = "\\[[0-9]{4}-[0-9]{2}-[0-9]{2}\\s[0-9]{2}:[0-9]{2}:[0-9]{2}\\.[0-9]{3}\\]\\s-\\s\\[.{5}\\]\\:\\s.*"

const levelInfo = "INFO"
const levelError = "ERROR"
const levelDebug = "DEBUG"
const messageStdout = "Hello World!"
const messageStderr = "Ups  ...Something it's wrong!"
const messageWithArgs = "Some message 'one', 'two' and 'three'"
const messagePrintMe = "Print me!"
const levelEnvironment = "GOUT_LEVEL"

func matchFormat(m string) bool {
	ok, _ := regexp.MatchString(outformat, m)
	return ok
}

func matchLevelAndMessage(f, l string, m string) bool {
	ml, _ := regexp.MatchString(fmt.Sprintf("\\[%-5s\\]", l), f)
	mm, _ := regexp.MatchString(fmt.Sprintf("\\:\\s%s", m), f)
	ok, _ := regexp.MatchString(outformat, f)
	return ml && mm && ok
}

func TestPrintStdout(t *testing.T) {
	c := out.Command{Level: out.LevelInfo{}, Message: "Hello {}", Placeholders: []string{"World!"}}
	m, s := Print(c)
	if s != os.Stdout {
		t.Error("The message must be printed to the STDOUT")
		t.Fail()
	}
	if !matchLevelAndMessage(m, levelInfo, messageStdout) {
		t.Error(fmt.Sprintf("The message must be equal to '[%-5s]: %s", levelInfo, messageStdout))
		t.Fail()
	}
}

func TestPrintStderr(t *testing.T) {
	c := out.Command{Level: out.LevelError{}, Message: "Ups {}", Placeholders: []string{" ...Something it's wrong!"}}
	m, s := Print(c)
	if s != os.Stderr {
		t.Error("The message must be printed to the STDERR")
		t.Fail()
	}
	if !matchLevelAndMessage(m, levelError, messageStderr) {
		t.Error(fmt.Sprintf("The message must be equal to '[%-5s]: %s", levelError, messageStderr))
		t.Fail()
	}
}

func TestPrintWithArgs(t *testing.T) {
	a := []string{"-d", "-m", "\"Some message '{}', '{}' and '{}'\"", "-p", "\"one\"", "-p", "two", "-p", "three"}
	m, s := PrintWithArgs(a)
	if s != os.Stdout {
		t.Error("The message must be printed to the STDOUT")
		t.Fail()
	}
	if !matchLevelAndMessage(m, levelDebug, messageWithArgs) {
		t.Error(fmt.Sprintf("The message must be equal to '[%-5s]: %s'", levelDebug, messageWithArgs))
		t.Fail()
	}
}

func TestMustNotBePrinted(t *testing.T) {
	defer os.Unsetenv(levelEnvironment)
	os.Setenv(levelEnvironment, levelError)
	a := []string{"-d", "-m", messagePrintMe}
	m, s := PrintWithArgs(a)
	if s != nil || m != "" {
		t.Error("The message mustn't be printed")
		t.Fail()
	}
}

func TestMustBePrinted(t *testing.T) {
	defer os.Unsetenv(levelEnvironment)
	os.Setenv(levelEnvironment, levelDebug)
	a := []string{"-d", "-m", "\"Print me!\""}
	m, s := PrintWithArgs(a)
	if s != os.Stdout {
		t.Error("The message must be printed to the STDOUT")
		t.Fail()
	}
	if !matchLevelAndMessage(m, levelDebug, messagePrintMe) {
		t.Error(fmt.Sprintf("The message must be equal to '[%-5s]: %s'", levelDebug, messagePrintMe))
		t.Fail()
	}
}

func TestGetVersion(t *testing.T) {
	a := []string{"--version"}
	m, s := PrintWithArgs(a)

	if s != os.Stdout {
		t.Error("The message must be printed to the STDOUT")
		t.Fail()
	}

	if !strings.Contains(m, VERSION) {
		t.Error(fmt.Sprintf("Version must be %s", VERSION))
		t.Fail()
	}
}

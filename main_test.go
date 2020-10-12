package main

import (
	"os"
	"testing"

	"github.com/pascencio/goutput/out"
)

func TestPrintStdout(t *testing.T) {
	c := out.Command{Level: out.LevelInfo{}, Message: "Hello {}", Placeholders: []string{"World!"}}
	m, s := Print(c)
	if s != os.Stdout {
		t.Error("The message must be printed to the STDOUT")
		t.Fail()
	}
	if m != "[INFO]: Hello World!\n" {
		t.Error("The message must be equal to '[INFO]: Hello World!\\n'")
		t.Fail()
	}
}

func TestPrintStderr(t *testing.T) {
	c := out.Command{Level: out.LevelError{}, Message: "Ups {}", Placeholders: []string{" ...Something it's wrong!"}}
	m, s := Print(c)
	if s != os.Stderr {
		t.Error("The message must be printed to the ERROR")
		t.Fail()
	}
	if m != "[ERROR]: Ups  ...Something it's wrong!\n" {
		t.Error("The message must be equal to '[ERROR]: Ups  ...Something it's wrong!\\n'")
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
	if m != "[DEBUG]: Some message 'one', 'two' and 'three'\n" {
		t.Error("The message must be equal to '[DEBUG]: Some message 'one', 'two' and 'three'\\n'")
		t.Fail()
	}
}

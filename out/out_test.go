package out

import (
	"os"
	"testing"
)

func TestCodeToLevel(t *testing.T) {
	debug := LevelDebug{}.GetValue()
	l, e := ArgumentToLevel("-d")
	if !e {
		t.Error("Argument DEBUG must exists")
		t.Fail()
	}
	if l.GetValue() != debug {
		t.Error("Argument must be have a value equal to DEBUG")
		t.Fail()
	}
	l, e = ArgumentToLevel("-x")
	if e {
		t.Error("Argument -x mustn't exists")
		t.Fail()
	}
}

func TestStringToLevel(t *testing.T) {
	info := LevelInfo{}.GetValue()
	error := LevelError{}.GetValue()
	l, e := StringToLevel("INFO")
	if !e {
		t.Error("Level INFO must exists")
		t.Fail()
	}
	if l.GetValue() != info {
		t.Error("Level must be have a value equal to INFO")
		t.Fail()
	}

	l, e = StringToLevel("ERROR")
	if !e {
		t.Error("Level ERROR must exists")
		t.Fail()
	}
	if l.GetValue() != error {
		t.Error("Level must be have a value equal to ERROR")
		t.Fail()
	}

	l, e = StringToLevel("DUMMY")

	if e {
		t.Error("Level DUMMY mustn't exists")
		t.Fail()
	}

}

func TestPrintStdout(t *testing.T) {
	c := Command{Level: LevelInfo{}, Message: "Hello {}", Placeholders: []string{"World!"}}
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
	c := Command{Level: LevelError{}, Message: "Ups {}", Placeholders: []string{" ...Something it's wrong!"}}
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

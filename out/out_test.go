package out

import (
	"testing"
)

func TestCodeToLevel(t *testing.T) {
	emptyArray := []string{}
	debug := LevelDebug{}.GetValue()
	l, e := ArgumentToLevel(map[string][]string{
		"-d": emptyArray,
	})
	if !e {
		t.Error("Argument DEBUG must exists")
		t.Fail()
	}
	if l.GetValue() != debug {
		t.Error("Argument must be have a value equal to DEBUG")
		t.Fail()
	}
	l, e = ArgumentToLevel(map[string][]string{
		"-x": emptyArray,
	})
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

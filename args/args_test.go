package args

import "testing"

func TestGetArgWithValue(t *testing.T) {
	a := []string{"-p", "dummy"}
	o := ParseArgs(a, "-p")
	v, ok := o["-p"]
	if !ok {
		t.Error("Map must contains value for -p argument")
		t.Fail()
	} else if len(v) != 1 {
		t.Error("Array inside map must one element")
		t.Fail()
	} else if v[0] != "dummy" {
		t.Error("First element in array must be equals to dummy")
		t.Fail()
	}
	print(o)
}

func TestGetArgWithoutValue(t *testing.T) {
	a := []string{"-p"}
	o := ParseArgs(a, "-p")
	v, ok := o["-p"]
	if !ok {
		t.Error("Map must contains value for -p argument")
		t.Fail()
	}
	if len(v) == 1 && v[0] != "" {
		t.Error("Array inside map must by empty")
		t.Fail()
	}
}

func TestGetWithManyArgs(t *testing.T) {
	a := []string{"-d", "-m", "Some message", "-p", "one", "-p", "two", "-p", "three"}
	o := ParseArgs(a, "-d", "-m", "-p")
	d, ok := o["-d"]
	if !ok {
		t.Error("Map must have level argument")
		t.Fail()
	}
	if len(d) == 1 && d[0] != "" {
		t.Error("Level debug mustn't have value")
		t.Fail()
	}
	m, ok := o["-m"]
	if !ok {
		t.Error("Map must have message argument")
		t.Fail()
	}
	if len(m) == 1 && m[0] == "" {
		t.Error("Message must have a value")
		t.Fail()
	}
	p, ok := o["-p"]
	if !ok {
		t.Error("Map must have placeholder argument")
		t.Fail()
	}
	if len(p) != 3 {
		t.Error("Placeholder must have three values")
		t.Fail()
	}
	if p[0] != "one" || p[1] != "two" || p[2] != "three" {
		t.Error("Placeholder must have the follwing values: 'one', 'two' and 'three'")
		t.Fail()
	}
}

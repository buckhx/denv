package denv

import (
	"testing"
)

func TestFreeze(t *testing.T) {
	//Push("https://github.com/buckhx/test-denv", "master")
}

/*
func TestFreeze(t *testing.T) {
	d := NewDenv("test-freeze")
	pkg := freeze("test-pkg")
	d.remove()
	thaw(pkg)
	d, _ = GetDenv("test-freeze")
	if d == nil {
		t.Errorf("Denv was not properly thawed")
	}
}
*/

/*
func TestPushPull(t *testing.T) {
	d := NewDenv("test-freeze")
	cases := []struct {
		in   string
		want *Denv
	}{
		{"", nil},
		{"test", d},
		{"notexist", nil},
	}
	for _, c := range cases {
		//fmt.Printf("Testing Activate(%s)\n", c.in)
		got, err := Activate(c.in)
		if c.want == nil && got != nil && err != nil {
			t.Errorf("Activate(%q) returned an error, but not nil Denv")
		} else if err != nil && c.want != nil {
			t.Errorf("Wanted something, but threw an error %s", err)
		} else if !reflect.DeepEqual(c.want, got) {
			t.Errorf("Activate(%q) wanted %s, got %s", c.in, c.want.ToString(), got.ToString())
		} else if got != nil {
			got.remove()
		}
	}
	d = NewDenv("test")
	denv, err := Activate("test")
	if err != nil {
		t.Errorf("Error creating Denv, %s", err)
	}
	if Info.Current != denv {
		t.Errorf("Activate(test) did not properly assign Info.Current, %s", Info.ToString())
	}
	d.remove()
}
*/

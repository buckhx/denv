package api

import (
	"fmt"
	"testing"
	"reflect"
)

func TestActivate(t *testing.T) {
	d := NewDenv("test")
	cases := []struct {
		in string
		want *Denv
	}{
		{"", nil},
		{"test", d},
		{"notexist", nil},
	}
	for _, c := range cases {
		fmt.Printf("Testing Activate(%s)\n", c.in)
		got, err := Activate(c.in)
		if c.want == nil && got != nil && err != nil {
			t.Errorf("Activate(%q) returned an error, but not nil Denv")
		} else if err != nil && c.want != nil {
			t.Errorf("Wanted something, but threw an error %s", err)
		} else if !reflect.DeepEqual(c.want, got) {
			t.Errorf("Activate(%q) wanted %s, got %s", c.in, c.want.ToString(), got.ToString())
		}
	}
	denv, err := Activate("test")
	if err != nil {
		t.Errorf("Error creating Denv, %s", err)
	}
	if Info.Current != denv {
		t.Errorf("Activate(test) did not properly assign Info.Current, %s", Info.ToString())
	}
}

func TestDeactivate(t *testing.T) {
	NewDenv("test") // Forces creation if not there
	active, err := Activate("test")
	if err != nil {
		t.Errorf("Could not Activate(test), %s", err)
	}
	deactive := Deactivate()
	if active != deactive {
		t.Errorf("Deactivate() reaturned different denv, %p, %p", &active, &deactive)
	}
	if Info.Current != nil {
		t.Errorf("Deactivate() did not clear Info.Current, %s", Info.ToString())
	}
}

func TestList(t *testing.T) {
	// Maybe implement a set
	denvs := map[*Denv]bool {
		NewDenv("d1"): true,
		NewDenv("d2"): true,
	}
	ls := List()
	set := make(map[string]int)
	for d, _ := range ls {
		set[d.Path] += 1
	}	
	for d, _ := range denvs {
		if count, found := set[d.Path]; !found {
			t.Errorf("Denvs not subset of List()\nDenvs %s\nList() %s", denvs, ls)
		} else if count < 1 {
			t.Errorf("Denvs not subset of List()\nDenvs %s\nList() %s", denvs, ls)
		} else {
			set[d.Path] = count - 1
		}
	}
}

func TestWhich(t *testing.T) {
	NewDenv("test")
	d, _ := Activate("test")
	if Which() != d {
		t.Errorf("Which() did not return active denv, %s, %s", d, Which())
	}
	d = Deactivate()
	if Which() != nil {
		t.Errorf("Which() did not return nil on deactivate, %s, %s", d, Which())
	}

}

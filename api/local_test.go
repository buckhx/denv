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
}

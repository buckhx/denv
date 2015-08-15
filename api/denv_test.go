package api

import (
	"io/ioutil"
	"os"
	pathlib "path"
	"testing"
)

func TestIgnore(t *testing.T) {
	var d *Denv
	d = NewDenv("test")
	err := os.RemoveAll(d.Path + "/*")
	check(err)
	patterns := []byte("test\n*.test")
	err = ioutil.WriteFile(pathlib.Join(d.Path, Settings.IgnoreFile), patterns, 0644)
	check(err)
	d = NewDenv("test")
	cases := []struct {
		in      string
		ignored bool
	}{
		{"", false},
		{"test", true},
		{"hey.test", true},
		{"test.txt", false},
	}
	for _, c := range cases {
		ignored := d.Ignored(c.in)
		if ignored != c.ignored {
			t.Errorf("Ignored(%q) did not ignore", c.in)
		}
	}
}

package api

import (
	"io/ioutil"
	"os"
	"os/user"
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
	usr, _ := user.Current()
	home := usr.HomeDir
	cases := []struct {
		in      string
		ignored bool
	}{
		{"", false},
		{"test", false},
		{pathlib.Join(home, "test"), true},
		{"hey.test", false},
		{"test.txt", false},
		{pathlib.Join(home, "hey.test"), true},
		{pathlib.Join(home, "test.txt"), false},
	}
	for _, c := range cases {
		ignored := d.Ignored(c.in)
		if ignored != c.ignored {
			t.Errorf("Ignored(%q) did not ignore", c.in)
		}
	}
	d.remove()
	d = NewDenv("test")
	cases = []struct {
		in      string
		ignored bool
	}{
		{".denv", false},
		{pathlib.Join(usr.HomeDir, ".denv"), true},
		{pathlib.Join(usr.HomeDir, ".bash_history"), true},
		{".bash_history", false},
	}
	for _, c := range cases {
		ignored := d.Ignored(c.in)
		if ignored != c.ignored {
			t.Errorf("Ignored(%q) did not ignore", c.in)
		}
	}
	d.remove()
}

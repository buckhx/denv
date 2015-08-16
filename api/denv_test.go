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
	patterns := []byte(".test\n*.test")
	err = ioutil.WriteFile(pathlib.Join(d.Path, Settings.IgnoreFile), patterns, 0644)
	check(err)
	d = NewDenv("test")
	usr, _ := user.Current()
	home := usr.HomeDir
	cases := []struct {
		in      string
		ignored bool
	}{
		{"", true},
		{".test", true},
		{pathlib.Join(home, ".test"), true},
		{pathlib.Join(home, "test"), true},
		{pathlib.Join(home, ".legit"), false},
		{"hey.test", true},
		{"test.txt", true},
		{pathlib.Join(home, "hey.test"), true},
		{pathlib.Join(home, ".hey.test"), true},
		{pathlib.Join(home, "test.txt"), true},
		{pathlib.Join(home, ".legit.txt"), false},
	}
	for _, c := range cases {
		ignored := d.IsIgnored(c.in)
		if ignored != c.ignored {
			t.Errorf("Ignored(%q) did not ignore", c.in)
		}
	}
	d.remove()
	d = NewDenv("test")
	//This are more for default ignores
	cases = []struct {
		in      string
		ignored bool
	}{
		{".denv", true},
		{pathlib.Join(usr.HomeDir, ".denv"), true},
		{pathlib.Join(usr.HomeDir, ".bash_history"), true},
		{pathlib.Join(usr.HomeDir, ".viminfo"), true},
		{pathlib.Join(usr.HomeDir, ".legit"), false},
		{".bash_history", true},
	}
	for _, c := range cases {
		ignored := d.IsIgnored(c.in)
		if ignored != c.ignored {
			t.Errorf("Ignored(%q) did not ignore", c.in)
		}
	}
	d.remove()
}

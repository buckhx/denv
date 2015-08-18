package api

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"os"
	"os/user"
	pathlib "path"
	"reflect"
	"testing"
)

func TestActivate(t *testing.T) {
	d := NewDenv("test")
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
	denv.remove()
}

func TestDeactivate(t *testing.T) {
	usr, _ := user.Current()
	home := usr.HomeDir
	testFile := pathlib.Join(home, ".test-deactivate.txt")
	check(ioutil.WriteFile(testFile, []byte("before"), 0664))
	Snapshot("test-deactivate")
	active, err := Activate("test-deactivate")
	if err != nil {
		t.Errorf("Could not Activate(test), %s", err)
	}
	check(ioutil.WriteFile(testFile, []byte("after"), 0664))
	deactive := Deactivate()
	if active != deactive {
		t.Errorf("Deactivate() reaturned different denv, %p, %p", &active, &deactive)
	}
	if Info.Current != nil {
		t.Errorf("Deactivate() did not clear Info.Current, %s", Info.ToString())
	}
	contents, err := ioutil.ReadFile(testFile)
	check(err)
	if string(contents) != "before" {
		t.Errorf("Deactivate() did not correctly restore the home directory")
	}
	os.Remove(testFile)
	active.remove()
}

func TestList(t *testing.T) {
	// Maybe implement a set
	denvs := map[*Denv]bool{
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
		d.remove()
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
	d.remove()
}

func TestSnapshot(t *testing.T) {
	usr, _ := user.Current()
	home := usr.HomeDir
	// could makke thse tempfiles
	testFile := pathlib.Join(home, ".test.txt")
	testDir := pathlib.Join(home, ".test")
	testDirFile := pathlib.Join(home, ".test/testdir.txt")
	err := os.MkdirAll(testDir, 0744)
	check(err)
	err = ioutil.WriteFile(testFile, []byte("derp"), 0664)
	check(err)
	err = ioutil.WriteFile(testDirFile, []byte("derp"), 0664)
	check(err)
	checks := map[string]bool{
		testFile:    false,
		testDir:     false,
		testDirFile: false,
	}
	d := Snapshot("test-snapshot")
	if d == nil {
		t.Errorf("Snapshot did not return Denv")
	}
	included, _ := d.Files()
	for _, path := range included {
		//t.Logf("Base path: %q, test: %q\n", pathlib.Base(path), pathlib.Base(testFile))
		if pathlib.Base(path) == pathlib.Base(testFile) {
			if fileCompare(testFile, path) == true {
				checks[testFile] = true
			}
		}
		if pathlib.Base(path) == pathlib.Base(testDir) {
			checks[testDir] = true
		}
		if pathlib.Base(path) == pathlib.Base(testDirFile) {
			if fileCompare(testDirFile, path) == true {
				checks[testDirFile] = true
			}
		}
	}
	for k, v := range checks {
		if v == false {
			t.Errorf("Snapshot did not persist %q correctly", k)
		}
	}
	os.Remove(testFile)
	os.Remove(testDir)
}

func fileCompare(first, second string) bool {
	fmt.Printf("fileCompare %q, %q\n", first, second)
	//Reads both into memory
	f1, err := ioutil.ReadFile(first)
	check(err)
	f2, err := ioutil.ReadFile(second)
	check(err)
	return bytes.Equal(f1, f2)
}

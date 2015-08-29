package denv

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestIgnore(t *testing.T) {
	var d *Denv
	d = NewDenv("test-ignore")
	err := os.RemoveAll(d.Path + "/*")
	check(err)
	patterns := []byte(".test\n*.test")
	err = ioutil.WriteFile(d.expandPath(Settings.IgnoreFile), patterns, 0644)
	d.LoadIgnore()
	check(err)
	cases := []struct {
		in      string
		ignored bool
	}{
		{"", true},
		{".test", true},
		{d.expandPath(".test"), true},
		{d.expandPath("test"), true},
		{d.expandPath(".legit"), false},
		{"hey.test", true},
		{"test.txt", true},
		{d.expandPath("hey.test"), true},
		{d.expandPath(".hey.test"), true},
		{d.expandPath("test.txt"), true},
		{d.expandPath(".legit.txt"), false},
	}
	for _, c := range cases {
		ignored := d.IsIgnored(c.in)
		if ignored != c.ignored {
			t.Errorf("IsIgnored(%q) != %t", c.in, c.ignored)
		}
	}
	d.remove()
	d = NewDenv("test-ignore")
	//This are more for default ignores
	cases = []struct {
		in      string
		ignored bool
	}{
		{".denv", true},
		{d.expandPath(".denv"), true},
		{d.expandPath(".denvignore"), true},
		{d.expandPath(".bash_history"), true},
		{d.expandPath(".viminfo"), true},
		{d.expandPath(".legit"), false},
		{".bash_history", true},
	}
	for _, c := range cases {
		ignored := d.IsIgnored(c.in)
		if ignored != c.ignored {
			t.Errorf("IsIgnored(%q) != %t", c.in, c.ignored)
		}
	}
	d.remove()
}

func TestInclude(t *testing.T) {
	d := NewDenv("test-include")
	ioutil.WriteFile(d.expandPath(".test.txt"), []byte("derp"), 0644)
	in, ex, _ := d.Files()
	wantIn := []string{d.expandPath(".test.txt")}
	wantEx := []string{d.expandPath(".denvignore")}
	if !reflect.DeepEqual(in, wantIn) {
		t.Errorf("Included files did not match, Want: %q, Got: %q", wantIn, in)
	}
	if !reflect.DeepEqual(ex, wantEx) {
		t.Errorf("Ignored files did not match, Want: %q, Got: %q", wantEx, ex)
	}
	os.Remove(".test.txt")
	d.remove()
}

func TestMatchedFiles(t *testing.T) {
	from := NewDenv("test-matched-files-from")
	to := NewDenv("test-matched-files-to")
	//TODO make api for changing gitignore
	want := []string{to.expandPath(".test.txt")}
	ioutil.WriteFile(want[0], []byte("derp"), 0644)
	got, _, _ := from.MatchedFiles(to.Path)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("MatchedFiles(%q) != %q, got %q", to.Path, want, got)
	}
	from.remove()
	to.remove()
}

func TestExpandPath(t *testing.T) {
}

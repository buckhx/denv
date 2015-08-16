package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	pathlib "path"
	"path/filepath"
	"strings"

	"github.com/buckhx/pathutil"
)

type Denv struct {
	Path   string
	Ignore map[string]bool
}

// Get a Denv that has already been created
// Nil, err if it hasn't been created yet
func GetDenv(name string) (*Denv, error) {
	if len(name) < 1 {
		return nil, fmt.Errorf("Denv name can't be empty")
	}
	path := pathlib.Join(Settings.DenvHome, name)
	if !pathutil.Exists(path) {
		return nil, fmt.Errorf("Denv %s does not exist", name)
	}
	return NewDenv(name), nil
}

// Create a new Denv
// Bootstrap the creation with the folder and default denvignore
func NewDenv(name string) *Denv {
	d := new(Denv)
	d.Path = pathlib.Join(Settings.DenvHome, name)
	d.bootstrap()
	d.LoadIgnore()
	return d
}

// Paths of files in the Denv Definition
// Tells you which paths are currently included and ignored
func (d *Denv) Files() (included []string, ignored []string) {
	return d.MatchedFiles(d.Path)
}

// Check to see if a file path be ignored by this Denv
func (d *Denv) IsIgnored(path string) bool {
	for pattern, _ := range d.Ignore {
		ignored, err := filepath.Match(pattern, path)
		check(err)
		if ignored == true {
			return true
		}
	}
	return false
}

func (d *Denv) IsNotIgnored(path string) bool {
	return !d.IsIgnored(path)
}

// Loads the denvignore from disk
func (d *Denv) LoadIgnore() {
	d.Ignore = make(map[string]bool)
	if pathutil.Exists(d.ignoreFile()) == true {
		content, err := ioutil.ReadFile(d.ignoreFile())
		check(err)
		//TODO handle comments and stuff
		patterns := strings.Split(string(content), "\n")
		usr, _ := user.Current()
		for _, pattern := range patterns {
			path := pathlib.Join(usr.HomeDir, pattern)
			d.Ignore[path] = true
		}
	} else {
		fmt.Printf("Warning: Denv %s has no .denvignore file at %s, all hidden files will be managed\n", d.Name(), d.ignoreFile())
	}
}

// Given an arbitrary path, return which files would be included
// and which would be ignored
func (d *Denv) MatchedFiles(string path) (inluded []string, ignored []string) {
	for _, f := range path {
		if d.Ignored(f) {
			ignored = append(ignored, f)
		} else {
			included = append(included, f)
		}
	}
	return
}

// Name of the denv
func (d *Denv) Name() string {
	return pathlib.Base(d.Path)
}

// String representation of the denv
func (d *Denv) ToString() string {
	content, err := json.Marshal(d)
	check(err)
	return string(content)
}

// Create the denv path and denvignore if they don't already exist
func (d *Denv) bootstrap() {
	if !pathutil.Exists(d.Path) {
		err := os.MkdirAll(d.Path, 0744)
		check(err)
		err = ioutil.WriteFile(d.ignoreFile(), []byte(_default_denvignore), 0644)
		check(err)
	}
}

// Path to the actual ignore file
func (d *Denv) ignoreFile() string {
	return pathlib.Join(d.Path, Settings.IgnoreFile)
}

// Completely remove this denv from disk
func (d *Denv) remove() error {
	return os.RemoveAll(d.Path)
}

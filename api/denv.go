package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	pathlib "path"
	"path/filepath"
	"strings"

	"github.com/buckhx/pathutil"
)

type Denv struct {
	Path   string
	Ignore map[string]bool
}

func (d *Denv) Ignored(path string) bool {
	for pattern, _ := range d.Ignore {
		ignored, err := filepath.Match(pattern, path)
		if err != nil {
			panic(err)
		}
		if ignored == true {
			return true
		}
	}
	return false
}

func (d *Denv) Name() string {
	return pathlib.Base(d.Path)
}

func (d *Denv) ToString() string {
	content, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func (d *Denv) LoadIgnore() {
	d.Ignore = make(map[string]bool)
	path := pathlib.Join(d.Path, Settings.IgnoreFile)
	if pathutil.Exists(path) == true {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		//TODO handle comments and stuff
		patterns := strings.Split(string(content), "\n")
		for _, pattern := range patterns {
			d.Ignore[pattern] = true
		}
	} else {
		fmt.Printf("Warning: Denv %s has no .denvignore file at %s, all hidden files will be managed\n", d.Name(), path)
	}
}

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

func NewDenv(name string) *Denv {
	d := new(Denv)
	d.Path = pathlib.Join(Settings.DenvHome, name)
	if !pathutil.Exists(d.Path) {
		err := os.MkdirAll(d.Path, 0744)
		if err != nil {
			panic(err)
		}
		//TODO: Bootstrap .gitignore
	}
	d.LoadIgnore()
	return d
}

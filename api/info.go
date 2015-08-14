package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/buckhx/pathutil"
)

type DenvInfo struct {
	Current *Denv
	Path    string
	Ignore	map[*regexp.Regexp]bool
}

var Info DenvInfo

func (d *DenvInfo) Clear() {
	d.Current = nil
}

func (d *DenvInfo) Flush() {
	content := []byte(d.ToString())
	err := ioutil.WriteFile(d.Path, content, 0644)
	if err != nil {
		panic(err)
	}
}

func (d *DenvInfo) IsActive() bool {
	return d.Current != nil
}

func (d *DenvInfo) Load() {
	//TODO make sure that this is an available file
	content, err := ioutil.ReadFile(d.Path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &d)
	if err != nil {
		panic(err)
	}
	if pathutil.Exists(Settings.IgnoreFile) {
		content, err := ioutil.ReadFile(d.Path)
		if err != nil {
			panic(err)
		}
		patterns := strings.Split(string(content), "\n")
		for _, pattern := range patterns {
			d.Ignore[regexp.MustCompile(pattern)] = true
		}
	}
}

func (d *DenvInfo) ToString() string {
	content, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func bootstrap() error {
	//TODO: maybe this should live somewhere else
	// Create DENVHOME
	if !pathutil.Exists(Settings.DenvHome) {
		err := os.MkdirAll(Settings.DenvHome, 0744)
		if err != nil {
			return err
		}
	}
	// Warn DENVIGNORE
	if !pathutil.Exists(Settings.IgnoreFile) {
		fmt.Printf("Warning: No .denvignore file at %s, all hidden files will be managed\n", Settings.IgnoreFile)
	}
	return nil
}

func init() {
	bootstrap()
	path := Settings.InfoFile
	Info.Path = path
	if !pathutil.Exists(path) {
		Info.Flush()
	}
	Info.Load()
}

package denv

import (
	"encoding/json"
	"io/ioutil"
	"os"

	git "github.com/buckhx/gitlib"
	"github.com/buckhx/pathutil"
)

type DenvInfo struct {
	Current    *Denv
	Path       string
	Repository *git.Repository
}

var Info DenvInfo

func (d *DenvInfo) Clear() {
	d.Current = nil
}

func (d *DenvInfo) Flush() {
	content := []byte(d.ToString())
	err := ioutil.WriteFile(d.Path, content, 0644)
	check(err)
}

func (d *DenvInfo) IsActive() bool {
	return d.Current != nil
}

func (d *DenvInfo) Load() {
	//TODO make sure that this is an available file
	content, err := ioutil.ReadFile(d.Path)
	check(err)
	err = json.Unmarshal(content, &d)
	check(err)
}

func (d *DenvInfo) ToString() string {
	content, err := json.Marshal(d)
	check(err)
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
	if !git.IsRepository(Settings.DenvHome) {
		repo, err := git.NewRepository(Settings.DenvHome)
		check(err)
		repo.Init()
		repo.Exclude("/.*") // exclude hidden root files
	}
	return nil
}

func init() {
	bootstrap()
	path := Settings.InfoFile
	Info.Path = path
	repo, err := git.NewRepository(Settings.DenvHome)
	check(err)
	Info.Repository = repo
	if !pathutil.Exists(path) {
		Info.Flush()
	}
	Info.Load()
}

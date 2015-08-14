package api

import (
	"encoding/json"
	"io/ioutil"
	"os"
	pathlib "path"

	"github.com/buckhx/pathutil"
)

type DenvInfo struct {
	Current *Denv
	Path    string
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
}

func (d *DenvInfo) ToString() string {
	content, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func init() {
	path := pathlib.Join(Settings.Path, Settings.InfoFile)
	Info.Path = path
	if !pathutil.Exists(Settings.Path) {
		err := os.MkdirAll(Settings.Path, 0744)
		if err != nil {
			panic(err)
		}
	}
	if !pathutil.Exists(path) {
		Info.Flush()
	}
	Info.Load()
}
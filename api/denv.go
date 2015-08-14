package api

import (
	"encoding/json"
	"fmt"
	"os"
	pathlib "path"

	"github.com/buckhx/pathutil"
)

type Denv struct {
	Path string
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

//TODO: CreateDenv (flush it to disk)

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
	}
	return d
}

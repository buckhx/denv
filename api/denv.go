package api

import (
	"fmt"
	pathlib "path"
)

type Denv struct {
	Path string
}

func (d *Denv) Name() string {
	return pathlib.Base(d.Path)
}

//TODO: CreateDenv (flush it to disk)

func GetDenv(name string) (*Denv, error) {
	path := pathlib.Join(Settings.Path, name)
	if !pathExists(path) {
		return nil, fmt.Errorf("Denv %s does not exist", name)
	}
	return NewDenv(name), nil
}

func NewDenv(name string) *Denv {
	d := new(Denv)
	d.Path = pathlib.Join(Settings.Path, name)
	return d
}

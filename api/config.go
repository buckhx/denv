package api

import (
	"io/ioutil"
	"os/user"
	pathlib "path"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Path     string
	InfoFile string
}

var Settings Config

func init() {
	path := "./settings.yml"
	config, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(config, &Settings)
	if err != nil {
		panic(err)
	}
	Settings.Path = pathExpand(Settings.Path)
}

func pathExpand(path string) string {
	if path[:2] == "~/" {
		usr, _ := user.Current()
		home := usr.HomeDir
		path = pathlib.Join(home, path[2:])
	}
	return path
}

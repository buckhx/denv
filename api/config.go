package api

import (
	"io/ioutil"
	pathlib "path"

	"gopkg.in/yaml.v2"
	"github.com/buckhx/pathutil"
)

type Config struct {
	Path     string
	InfoFile string
}

var Settings Config

func init() {
	path := "settings.yml"
	if !pathutil.Exists(path) {
		// for tests to reference correct settings
		path = pathlib.Join("..", path)
	}
	config, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(config, &Settings)
	if err != nil {
		panic(err)
	}
	Settings.Path = pathutil.Expand(Settings.Path)
}

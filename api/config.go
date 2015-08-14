package api

import (
	"io/ioutil"
	pathlib "path"

	"gopkg.in/yaml.v2"
	"github.com/buckhx/pathutil"
)

type Config struct {
	DenvHome	string
	IgnoreFile	string
	InfoFile string
}

var Settings Config

func init() {
	// TODO: create a settings lib
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
	if len(Settings.DenvHome) < 1 {
		panic("Missing DenvHome setting")
	}
	Settings.DenvHome = pathutil.Expand(Settings.DenvHome)
	if len(Settings.InfoFile) < 1 {
		panic("Missing InfoFile setting")
	}
	Settings.InfoFile = pathlib.Join(Settings.DenvHome, Settings.InfoFile)
	if len(Settings.IgnoreFile) < 1 {
		panic("Missing IgnoreFile setting")
	}
	Settings.IgnoreFile = pathlib.Join(Settings.DenvHome, Settings.IgnoreFile)
}

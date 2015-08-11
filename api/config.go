package api

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)


type Config struct {
	Denv struct {
		Path string
	}
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
}

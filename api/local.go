package api

import (
	"os"
)

func Activate(env string) []string {
	return []string{}
}

func Bootstrap() []string {
	if _, err := os.Stat(Settings.Denv.Path); os.IsNotExist(err) {
		err = os.MkdirAll(Settings.Denv.Path, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		panic("Dir already exists")
	}
	return []string{}
}

func List() []string {
	return []string{"one", "two"}
}

func Which() []string {
	return []string{"WHICHI"}
}

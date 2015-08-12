package api

import (
	"os"
	"path/filepath"
)

func Activate(env string) []string {
	return []string{}
}

func Bootstrap() []string {
	if pathExists(Settings.Denv.Path) {
		_ = os.MkdirAll(Settings.Denv.Path, 0744)
	} else {
		return []string{"Already Bootstrapped " + Settings.Denv.Path}
	}
	return []string{"Created " + Settings.Denv.Path}
}

func List() []string {
	//TODO Check is Settings.Denv.Path exists
	denvs := []string{}
	err := filepath.Walk(Settings.Denv.Path, func(path string, file os.FileInfo, err error) error {
		if file.IsDir() && path != Settings.Denv.Path {
			denvs = append(denvs, file.Name())
		}
		return err
	})
	if err != nil {
		panic(err)
	}
	return denvs
}

func Which() []string {
	return []string{"WHICHI"}
}

func pathExists(path string) bool {
	_, err := os.Stat(Settings.Denv.Path)
	return os.IsNotExist(err)
}

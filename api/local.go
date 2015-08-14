package api

import (
	"os"
	"path/filepath"

	"github.com/buckhx/pathutil"
)

func Activate(env string) (*Denv, error) {
	denv, err := GetDenv(env)
	if err != nil {
		return nil, err
	}
	Info.Current = denv
	Info.Flush()
	return denv, nil
}

//Boostrap the denv envirnoment from the settings
//If it was already bootstrapped, nothing happens
//Returns the path of the denv setup
func Bootstrap() string {
	if !pathutil.Exists(Settings.Path) {
		_ = os.MkdirAll(Settings.Path, 0744)
	}
	return Settings.Path
}

//Deactivate the current denv and restore it to the state
//before denv was active. Returns the name of the deactivated denv.
//Empty string if there was no denv to deactivate
func Deactivate() *Denv {
	denv := Info.Current
	if Info.IsActive() {
		Info.Clear()
		Info.Flush()
	}
	return denv
}

func List() map[*Denv]bool {
	//TODO Check is Settings.Denv.Path exists
	denvs := make(map[*Denv]bool)
	//TODO decide if this logic should be moved to DenvInfo
	err := filepath.Walk(Settings.Path, func(path string, file os.FileInfo, err error) error {
		if file.IsDir() && path != Settings.Path {
			denvs[NewDenv(file.Name())] = true
		}
		return err
	})
	if err != nil {
		panic(err)
	}
	return denvs
}

func Which() *Denv {
	return Info.Current
}

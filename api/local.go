package api

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	pathlib "path"
	"path/filepath"
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
	err := filepath.Walk(Settings.DenvHome, func(path string, file os.FileInfo, err error) error {
		if file.IsDir() && path != Settings.DenvHome {
			denvs[NewDenv(file.Name())] = true
		}
		return err
	})
	check(err)
	return denvs
}

func Which() *Denv {
	return Info.Current
}

func Snapshot(name string) *Denv {
	d, _ := GetDenv(name)
	if d == nil {
		fmt.Printf("Denv didn't exist, bootstrapping %s\n", name)
		d = NewDenv(name)
	}
	usr, _ := user.Current()
	included, _ := d.MatchedFiles(usr.HomeDir)
	for _, src := range included {
		dst := d.expandPath(pathlib.Base(src))
		err := cp(src, dst)
		check(err)
	}
	return d
}

func cp(src, dst string) error {
	fmt.Printf("cp -rf %s %s\n", src, dst)
	cmd := exec.Command("cp", "-rf", src, dst)
	return cmd.Run()
}

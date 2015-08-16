package api

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	pathlib "path"
	"path/filepath"
	"strings"
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
	var denvfiles []string
	for _, f := range homeFiles() {
		if d.IsDenvFile(f) {
			denvfiles = append(denvfiles, f)
		}
	}
	for _, src := range denvfiles {
		dst := pathlib.Join(d.Path, pathlib.Base(src))
		err := cp(src, dst)
		check(err)
	}
	return d

}

func cp(src, dst string) error {
	cmd := exec.Command("cp", "-rf", src, dst)
	return cmd.Run()
}

func homeFiles() []string {
	homefiles := []string{}
	usr, _ := user.Current()
	err := filepath.Walk(usr.HomeDir, func(path string, file os.FileInfo, err error) error {
		if strings.HasPrefix(file.Name(), ".") && path == pathlib.Join(usr.HomeDir, file.Name()) {
			homefiles = append(homefiles, path)
		}
		return err
	})
	check(err)
	return homefiles
}

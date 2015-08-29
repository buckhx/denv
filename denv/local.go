package denv

import (
	"fmt"
	"os"
	"os/exec"
	pathlib "path"
	"path/filepath"
	"strings"
)

func Activate(env string) (*Denv, error) {
	denv, err := GetDenv(env)
	if err != nil {
		return nil, err
	}
	if !Info.IsActive() {
		stash(denv) // only stash the homedir if no denv is active
	} else {
		Info.Current.Exit()
	}
	denv.Enter()
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
		Info.Current.Exit()
		restore, _ := GetDenv(Settings.RestoreDenv)
		if restore == nil {
			fmt.Printf("WARNING: There was no RestoreDenv at %s, something looks fishy...\n", Settings.RestoreDenv)
		} else {
			restore.Enter()
		}
		Info.Clear()
		Info.Flush()
	}
	return denv
}

// TODO: Make a ls denv -> files
func List() map[*Denv]bool {
	//TODO Check is Settings.Denv.Path exists
	denvs := make(map[*Denv]bool)
	//TODO decide if this logic should be moved to DenvInfo
	err := filepath.Walk(Settings.DenvHome, func(path string, file os.FileInfo, err error) error {
		if file.IsDir() && path != Settings.DenvHome {
			if !strings.HasPrefix(file.Name(), ".") {
				denvs[NewDenv(file.Name())] = true
			}
			return filepath.SkipDir // don't recurse
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
		//fmt.Printf("Denv didn't exist, bootstrapping %s\n", name)
		//TODO: only squash when -f flag is passed
		d = NewDenv(name)
	}
	included, _, _ := d.MatchedFiles(UserHome())
	for _, src := range included {
		//TODO: only copy root files and dirs
		dst := d.expandPath(pathlib.Base(src))
		err := fileCopy(src, dst)
		if err != nil {
			fmt.Printf("WARNING: Could not copy %s to %s, skipping...", src, dst)
		}
	}
	d.cleanGitSubmodules()
	return d
}

func stash(denv *Denv) {
	snap := NewDenv(Settings.RestoreDenv)
	snap.SetDenvIgnore(denv.ignoreFile())
	snap = Snapshot(Settings.RestoreDenv)
}

//TODO: move to a util
func fileCopy(src, dst string) error {
	//fmt.Printf("\tcp -rf %s %s\n", src, dst)
	cmd := exec.Command("cp", "-rf", src, dst)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("ERROR: cp -rf %s %s\n", src, dst)
	} else {
		return nil
	}
}

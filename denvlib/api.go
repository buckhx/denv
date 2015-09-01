/*
denvlib provides the mechanics for the denv client.
The denv client handles the hidden files in the users home directory
which are generally configurations for the users environment.

A simple example use case is when doing Python and Go development on the same box.
Python (PEP-8) wants 4 spaces for it's whereas Go likes actual tab characters.
Being able to quickly switch between these by activating a denv solves this iise.

Each denv manages a particular environment and can be activated/deactived.

Denv definitions can be pushed/pulled from a remote git server
*/
package denvlib

import (
	"fmt"
	"os"
	"os/exec"
	pathlib "path"
	"path/filepath"
	"strings"
)

//Activate a denv by replacing hidden files in the users home directoty with those in the denv definition.
//as well as executing the scripts .denvpre in the denv definition
//If one is not already active, the current state of the users home directory
//will be stashed and restored whenever the user deactivates.
//Returns the denv that was activated, nil if there was an error
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

//Deactivate the current denv and restore it to the state. Also executes the denvs .denvpost scripts.
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

//Gets a map of the denvs currently on the system
//The returned map is meanted to be used like a set
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

//Pull the contents of the remote/branch from the remote server onto the local system.
//If there are conflicts, they will need to be managed manually in ~/.denv
//This also makes denv scripts executable with chmod 744
func Pull(remote string, branch string) string {
	Info.Repository.SetRemote("denv", remote)
	Info.Repository.Fetch("denv")
	Info.Repository.Checkout("-b", branch)
	Info.Repository.Checkout(branch)
	Info.Repository.Pull("denv", branch)
	for d := range List() {
		_, _, scripts := d.Files()
		for _, script := range scripts {
			os.Chmod(script, 0744)
		}
	}
	return ""
}

//Push the current contents of ~/.denv to a remote git server at the specified branch.
//If there are issues, they will need to be resolved manually with git @ ~/.denv
func Push(remote string, branch string) string {
	Info.Repository.SetRemote("denv", remote)
	Info.Repository.Fetch("denv")
	Info.Repository.Checkout("-b", branch)
	Info.Repository.Checkout(branch)
	Info.Repository.Add(".")
	Info.Repository.Commit("freeze")
	Info.Repository.Pull("denv", branch)
	Info.Repository.Push("denv", branch)
	return ""
}

//Creates copies of the hidden files in the users home directory and puts them in a new denv definition
//in ~/.denv/name. The denv definition can be edited manually after. The default denvignore will be used.
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

// Returns the currently active Denv or nil
func Which() *Denv {
	return Info.Current
}

// save the current user home to RestoreDenv
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

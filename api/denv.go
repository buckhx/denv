package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	pathlib "path"
	"path/filepath"
	"strings"
	"sort"

	"github.com/buckhx/pathutil"
)

type Denv struct {
	Path   string
	Ignore map[string]bool
}

// Get a Denv that has already been created
// Nil, err if it hasn't been created yet
func GetDenv(name string) (*Denv, error) {
	if len(name) < 1 {
		return nil, fmt.Errorf("Denv name can't be empty")
	}
	path := pathlib.Join(Settings.DenvHome, name)
	if !pathutil.Exists(path) {
		return nil, fmt.Errorf("Denv %s does not exist", name)
	}
	return NewDenv(name), nil
}

// Create a new Denv
// Bootstrap the creation with the folder and default denvignore
func NewDenv(name string) *Denv {
	d := new(Denv)
	d.Path = pathlib.Join(Settings.DenvHome, name)
	d.bootstrap()
	d.LoadIgnore()
	return d
}

func (d *Denv) Enter() {
	included, _, scripts := d.Files()
	for _, script := range scripts {
		if strings.HasSuffix(script, Settings.PreScript) {
			run_script(script)
		}
	}
	for _, src := range included {
		dst := pathlib.Join(UserHome(), pathlib.Base(src))
		err := fileCopy(src, dst)
		if err != nil {
			fmt.Printf("WARNING: Could not copy %s to %s, skipping...", src, dst)
		}
	}
}

func (d *Denv) Exit() {
	_, _, scripts := d.Files()
	for _, script := range scripts {
		if strings.HasSuffix(script, Settings.PostScript) {
			run_script(script)
		}
	}
}

//TODO addd denv.AddFile(path)
// Paths of files in the Denv Definition
// Tells you which paths are currently included and ignored
func (d *Denv) Files() (included []string, ignored []string, scripts []string) {
	return d.MatchedFiles(d.Path)
}

// Check to see if a file path be ignored by this Denv
func (d *Denv) IsIgnored(path string) bool {
	//path must include denvpath and file is hidden
	if !strings.HasPrefix(path, d.Path+"/.") {
		//fmt.Printf("Path: %q, Prefix: %q\n", path, d.Path+"/.")
		return true
	}
	for pattern, _ := range d.Ignore {
		ignored, err := filepath.Match(pattern, path)
		//TODO support inverse and include patterns
		//if strings.HasPrefix(pattern, ".!") {
		//	ignored = !ignored
		//}
		//fmt.Printf("path: %q, pattern: %q, ignored: %t\n", path, pattern, ignored)
		check(err)
		if ignored == true {
			return true
		}
	}
	return false
}

func (d *Denv) IsScript(path string) bool {
	if strings.HasSuffix(path, Settings.PreScript) || strings.HasSuffix(path, Settings.PostScript) {
		return true
	} else {
		return false
	}
}

func (d *Denv) IsDenvFile(path string) bool {
	return !d.IsIgnored(path)
}

// Loads the denvignore from disk
func (d *Denv) LoadIgnore() {
	d.Ignore = make(map[string]bool)
	if pathutil.Exists(d.ignoreFile()) == true {
		content, err := ioutil.ReadFile(d.ignoreFile())
		check(err)
		//TODO handle comments and stuff
		patterns := strings.Split(string(content), "\n")
		for _, pattern := range patterns {
			path := d.expandPath(pattern)
			d.Ignore[path] = true
		}
	} else {
		fmt.Printf("Warning: Denv %s has no .denvignore file at %s, all hidden files will be managed\n", d.Name(), d.ignoreFile())
	}
}

// Given an arbitrary path, return which files would be included
// and which would be ignored. If path contains a folder, it will not
// be recursed into.
func (d *Denv) MatchedFiles(root string) (included []string, ignored []string, scripts []string) {
	err := filepath.Walk(root, func(path string, file os.FileInfo, err error) error {
		//chpath is created for when testing denvfiles against another dir
		chpath := strings.Replace(path, root, d.Path, 1)
		if path == root {
			return err // allows to recursively inspect root
		} else if d.IsScript(chpath) {
			scripts = append(scripts, path)
		} else if d.IsIgnored(chpath) {
			ignored = append(ignored, path)
		} else {
			included = append(included, path)
		}
		if file.IsDir() {
			return filepath.SkipDir
		}
		return err
	})
	sort.Strings(scripts)
	check(err)
	return
}

// Name of the denv
func (d *Denv) Name() string {
	return pathlib.Base(d.Path)
}

func (d *Denv) SetDenvIgnore(path string) {
	// TODO: assert is a .denvignore
	if path != d.ignoreFile() {
		err := fileCopy(path, d.ignoreFile())
		check(err)
	}
	d.LoadIgnore()
}

// String representation of the denv
func (d *Denv) ToString() string {
	content, err := json.Marshal(d)
	check(err)
	return string(content)
}

// Create the denv path and denvignore if they don't already exist
func (d *Denv) bootstrap() {
	if !pathutil.Exists(d.Path) {
		err := os.MkdirAll(d.Path, 0744)
		check(err)
		err = ioutil.WriteFile(d.ignoreFile(), []byte(_default_denvignore), 0644)
		check(err)
	}
}

// We want to track contents of internal git repos
// This should be done with Git Submodules, but for now
// We'll just clean out internal git repos
// TODO: Use submodules for tracking
func (d *Denv) cleanGitSubmodules() {
	err := filepath.Walk(d.Path, func(path string, file os.FileInfo, err error) error {
		if strings.HasPrefix(file.Name(), ".git") {
			err = os.RemoveAll(path)
		}
		return err
	})
	check(err)
}

func (d *Denv) expandPath(path string) string {
	return pathlib.Join(d.Path, path)
}

// Path to the actual ignore file
func (d *Denv) ignoreFile() string {
	return pathlib.Join(d.Path, Settings.IgnoreFile)
}

// Completely remove this denv from disk
func (d *Denv) remove() error {
	return os.RemoveAll(d.Path)
}

func operation(command string, args ...string) (string, string, error) {
	var stderr, stdout bytes.Buffer
	cmd := exec.Command(command, args...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func run_script(script string) {
	fmt.Printf("Executing %s...\n", script)
	stdout, stderr, err := operation(script)
	if err != nil {
		fmt.Println(err)
	}
	if len(stdout) > 0 {
		fmt.Printf(stdout)
	}
	if len(stderr) > 0 {
		fmt.Printf(stderr)
	}
}

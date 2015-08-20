package api

import (
	"fmt"
	"bytes"
	"os/exec"
	pathlib "path"
)

func Pull(remote string) string {
	//pkg := pull(remote)
	//pkg = decrypt(pkg, passphrase)
	//thaw(pkg)
	//return pkg.dirs()
	return ""
}

func Push() string {
	//pkg := freeze()
	//pkg = encrypt(pkg, passphrase)
	//push(pkg, remote)
	//return remote
	return ""
}

func freeze(name string) (pkg string) {
	var stderr bytes.Buffer
	pkg = pathlib.Join(Settings.Freezer, name)
	tar_gpg := fmt.Sprintf("tar czvpf - %s --exclude=\"\\./denv/\\.*\" | gpg --symmetric --cipher-algo aes256 -o %s", Settings.DenvHome, pkg)
	cmd := exec.Command("bash", "-c", tar_gpg)
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		err = fmt.Errorf("%s\n%s\n", stderr.String(), tar_gpg)
	}
	check(err)
	return
}

func thaw(pkg string) {
	var stderr bytes.Buffer
	untar_gpg := fmt.Sprintf("gpg -d %s | tar xzvf - -C %s", pkg, UserHome())
	cmd := exec.Command("bash", "-c", untar_gpg)
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		err = fmt.Errorf("%s\n%s\n", stderr.String(), untar_gpg)
	}
	check(err)
	//return thawed denvs
}

func compress(paths []string) string {
	//output := "denv.tar.gz.gpg"
	//cmd := fmt.Sprintf("tar czvpf - %s --exclude="\./denv/\.*" | gpg --symmetric --cipher-algo aes256 -o %s", Settings.DenvHome, output)
	return "" 

}

func decompress(path string) []string {
	//cmd := fmt.Sprintf("gpg -d %s | tar xzvf -C %s", path, Settings.DenvHome)
	return []string{}

}

func encrypt(pkg string, passphrase string) string {
	return ""

}

func decrypt(pkg string, passphrase string) string {
	return ""

}

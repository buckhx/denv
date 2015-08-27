package api

import (
	"bytes"
	"fmt"
	"os/exec"
	pathlib "path"
)

//TODO: change the returns
func Pull(remote string, branch string) string {
	Info.Repository.SetRemote("denv", remote)
	Info.Repository.Checkout("-b", branch)
	Info.Repository.Checkout(branch)
	Info.Repository.Pull("denv", branch)
	return ""
}

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

func freeze(name string) (pkg string) {
	var stderr bytes.Buffer
	pkg = pathlib.Join(Settings.Freezer, name)
	tar_gpg := fmt.Sprintf("tar czvpfC - %s %s --exclude=\"\\./denv/\\.*\" | gpg --symmetric --cipher-algo aes256 -o %s", UserHome(), pathlib.Base(Settings.DenvHome), pkg)
	cmd := exec.Command("bash", "-c", tar_gpg)
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		err = fmt.Errorf("%s\n%s\n", stderr.String(), tar_gpg)
	}
	fmt.Println(tar_gpg)
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
	fmt.Println(untar_gpg)
	check(err)
	//return thawed denvs
}

func compress(paths []string) string {
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

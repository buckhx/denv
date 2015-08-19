package api

import (
	"archive/tar"
	"bytes"
)

func Pull(remote string) string {
	pkg := pull(remote)
	pkg = decrypt(pkg, passphrase)
	thaw(pkg)
	return pkg.dirs()
}

func Push() string {
	pkg := freeze()
	pkg = encrypt(pkg, passphrase)
	push(pkg, remote)
	return remote
}

func freeze() string {
	//fuck it, just os.Exec this bad boy
	//only take denvs
}

func thaw() {
	//return thawed denvs
}

func compress(paths []string) path string {
// tar zcvf denv.tar.gz ./denv --exclude="\./denv/\.*"

}

func decompress(path string) paths []string{

}

func encrypt(pkg string, passphrase string) string {

}

func decrypt(pkg string, passphrase string) string {

}

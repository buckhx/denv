package api

import (
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

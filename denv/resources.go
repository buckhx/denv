package denv

const (
	settings_yml = `denvhome: ~/.denv
infofile: .denvinfo
ignorefile: .denvignore
restoredenv: .restore
prescript: .denvpre
postscript: .denvpost
`
	_default_denvignore = `.bash_history
.gnupg
.npm
.nvm
.rvm
.gimme
.cache
.gem
.bundle
.erlang.cookie
.git
.denv*
.ssh
.viminfo
`
	Version = "v0.0.6-a"
)

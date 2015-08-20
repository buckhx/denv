package api 

const (
settings_yml = `denvhome: ~/.denv
infofile: .denvinfo
ignorefile: .denvignore
restoredenv: .restore
freezer: .freezer
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
Version = "3e28a10"
)

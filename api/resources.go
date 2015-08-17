package api 

const (
settings_yml = `denvhome: ~/.denv
infofile: .denvinfo
ignorefile: .denvignore
`
_default_denvignore = `.bash_history
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
)

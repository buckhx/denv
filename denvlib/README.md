# denvelopment

## Contributions

TODO

## Development Notes

I still need to work through this part a bit

### Resources

Denv requires a few static assets for sensible defaults. In order to accomodate 
this as well as not requiring Denv to have to depend on static paths, resources
are included by generating them via the scripts/include.go routine. This should
be a pre-build step and in go >= 1.4 is included via `go generate`. Older versions
should run `go run scripts/include.go` before building to embed the assets. If
a resource is changed (settings.yml) without running include.go beforehand, the
changes will not take.

###DenvInfo

Maintains state

###DenvHome

Home dir for Denv to live in. Generally a hidden path in the users $HOME

###DenvIgnore

Do not keep file handles that match these patterns
Similar to a .gitignore file


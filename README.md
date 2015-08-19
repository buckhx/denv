# denv
[![Build Status](https://travis-ci.org/buckhx/denv.svg)](https://travis-ci.org/buckhx/denv)

Installs a way to manage your development environments

Really this is a just a sample go project that will let
you switch between your vim environments

## Installation

Not sure yet...

## Usage

Let's start out by listing the current denvs

    denv ls

Now let's turn one of these bad boys on

    denv activate mydenv

Actually, let's not

    denv deactivate

Really I want to save my current environment to a new denv

    denv snapshot newdenv

And push all my denvs to a remote server (passphrase required)

    denv push https://github.com/buckhx/somepath

Log on from a different VM or computer and pull what you just put up there

    denv pull https://github.com/buckhx/somepath

Back on the same page again!

    denv ls
    denv activate newdenv

## Development

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


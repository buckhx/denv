# denv
[![Build Status](https://travis-ci.org/buckhx/denv.svg)](https://travis-ci.org/buckhx/denv)

Installs a way to manage your development environments using the denv command



## Installation

Install denv into /usr/local/bin (assuming you have permissions to) like so

    curl -sSL https://raw.githubusercontent.com/buckhx/denv/master/scripts/install.py | python - /usr/local/bin

Or do it manually by going to the releases page and download the denv artifact

    https://github.com/buckhx/denv/releases/latest 


## Usage

Let's start out by listing the current denvs

    denv ls
    
You probably won't see anything, so why don't we download some denvs

    dev pull https://github.com/buckhx/template-denv.git
    
If you ls again, you should see some stuff. 
Check out your ~/.vimrc and see what it looks like before activation. 
Now let's turn one of these bad boys on.

    denv activate python

Take another peek at your ~/.vimrc and you'll notice it will be the one in the denv definition you downloaded. 
A keen eye will see that the new vimrc disables arrow keys and you're not into that. 
Let's revert to your old state by deactivating.

    denv deactivate

When you deactive, the state will be restored before there were ANY denvs active. 
So if you activate a few in a row before you explicitly deactivate, you'll be back to where you started. 
Now let's say I want to create a new denv defition from my current home directory.

    denv snapshot newdenv

If you inspect ~/.denv/newdenv, you'll notice that some of your files from you home have been copied over. 
Any file which does not match a pattern in ~/.denv/newdenv/.denvigore will be copied. 
Now you can do some cleaning up in the denv definition and push your new denv (along with any changes you made to the other ones) to a remote git server. 
All the denvs are managed via git, so feel free to create a fork of the template to push to.
If you have contributions to the template-denv, just shoot over a PR.

    denv push https://github.com/YOURNAMEHERE/template-denv.git

Log on from a different VM or computer and pull what you just put up there

    denv pull https://github.com/YOURNAMEHERE/template-denv.git

Back on the same page again!

    denv ls
    denv activate newdenv

## Development and Contributions

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


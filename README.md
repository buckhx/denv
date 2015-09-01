# denv
[![Build Status](https://travis-ci.org/buckhx/denv.svg)](https://travis-ci.org/buckhx/denv)

Installs a way to manage your development environments using the denv command

## Installation

Install denv into /usr/local/bin (assuming you have permissions to) like so

    curl -sSL https://raw.githubusercontent.com/buckhx/denv/master/scripts/install.py | python - /usr/local/bin
    
If you don't have permission at /usr/local/bin, try something like this where you extend your PATH

    BINDIR=~/.denv/.bin/
    mkdir -p $BINDIR
    export PATH=$PATH:$BINDIR
    curl -sSL https://raw.githubusercontent.com/buckhx/denv/master/scripts/install.py | python - $BINDIR

Or do it manually by going to the releases page and download the denv artifact https://github.com/buckhx/denv/releases/latest 


## Usage

Let's start out by listing the current denvs

    denv ls
    
You probably won't see anything, so why don't we download some denvs

    denv pull https://github.com/buckhx/template-denv.git
    
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

See README in https://github.com/buckhx/denv/tree/master/denvlib

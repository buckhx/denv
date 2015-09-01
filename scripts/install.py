#!usr/bin/e/nv python

'''
This script has not dependencies besides a conntection to github in order to run.
Therefore it can be copied from the repo and ran on any system

By default it will install to /usr/local/bin/denv. Change it by editing the INSTALL_LOCATION
'''
import json
import os
import sys
import urllib

# Change this if you don't have root to somewhere else on your path
if len(sys.argv) > 1:
    bindir = os.path.join(sys.argv[1], 'denv')
else:
    bindir = '/usr/local/bin/denv'
print "Installing denv into {0}...".format(bindir)

content = urllib.urlopen('https://api.github.com/repos/buckhx/denv/releases/latest').read()
release = json.loads(content)
link = [asset['browser_download_url'] for asset in release['assets'] if asset['name'] == 'denv'][0]
print "Downloading binary from " + link
urllib.urlretrieve(link, bindir)
os.chmod(bindir, 0755)
print "Installed denv at "+bindir

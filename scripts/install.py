#!/usr/bin/env python

'''
This script has not dependencies besides a conntection to github in order to run.
Therefore it can be copied from the repo and ran on any system

By default it will install to /usr/local/bin/denv. Change it by editing the INSTALL_LOCATION
'''

import json
import os
import urllib

# Change this if you don't have root to somewhere else on your path
INSTALL_LOCATION = '/usr/local/bin/denv'

content = urllib.urlopen('https://api.github.com/repos/buckhx/denv/releases/latest').read()
release = json.loads(content)
link = [asset['browser_download_url'] for asset in release['assets'] if asset['name'] == 'denv'][0]
urllib.urlretrieve(link, INSTALL_LOCATION)
os.chmod(INSTALL_LOCATION, 0111)

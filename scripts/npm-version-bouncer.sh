#!/bin/bash

# ---
# Description:
#   This will compare a package version against a current package version available
#   via npm.
#
#   You can either manually specify the package and version you're comparing,
#   or simply reference the package json of the package you are prospecting
#   for release.
#
# Usage: 
#   sh npm-version-bouncer ./path/to/package.json
#   sh npm-version-bouncer package-name version-number
# ---


PACKAGE_NAME=""
VERSION=""

if [ -z "$1" ]; then
    echo "node-version-bouncer [package.json] || [packagename] [version]"    
    exit 1
fi

PACKAGE_JSON=$(node -e "
    if ('$1'.substr('$1'.length - 12) === 'package.json') {
        try {
            console.log(!!JSON.parse(require('fs').readFileSync('$1', 'utf8'))) 
        } catch (err) {
            console.log('ERROR: Invalid package.json')
            process.exit(1)
        }
    }               
")

if [ "$PACKAGE_JSON" = true ]; then
    PACKAGE_NAME=$(node -e "console.log(JSON.parse(require('fs').readFileSync('$1', 'utf8')).name)")
    VERSION=$(node -e "console.log(JSON.parse(require('fs').readFileSync('$1', 'utf8')).version)")
else
    if [ -z "$2" ]; then
        echo "node-version-bouncer [package.json] || [packagename] [version]"    
        exit 1
    fi
    
    PACKAGE_NAME=$1
    VERSION=$2
fi

CURRENT_VERSION=$(npm view $PACKAGE_NAME version)

# Semver comparison taken and amended from:
# https://github.com/substack/semver-compare/blob/master/index.js
node -e "
    let exit = (result, code) => { console.log(result); process.exit(code) }
    let pa = '$VERSION'.split('.')
    let pb = '$CURRENT_VERSION'.split('.')
    if (pa.length != 3 || pb.length != 3) exit('ERROR: Invalid version format', 1)
    for (let i = 0; i < 3; i++) {
        let na = Number(pa[i])
        let nb = Number(pb[i])
        if (na > nb) exit('SUCCESS: Version $VERSION is a valid package version', 0)
        if (nb > na) exit('ERROR: Version $VERSION is lower than the latest NPM package version', 1)
        if (!isNaN(na) && isNaN(nb)) exit('SUCCESS: Version $VERSION is a valid package version', 0)
        if (isNaN(na) && !isNaN(nb)) exit('ERROR: Version $VERSION is lower than the latest NPM package version', 1)
    }
    exit('ERROR: Version $VERSION is the same as the latest NPM package version', 1)
"

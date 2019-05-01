#!/bin/sh
PLATFORM=$1
DEST=/usr/local/lib/node-sass

if [ -z $1 ]; then
    PLATFORM=$(node -e "console.log(require('os').platform())")
fi

for URL in $(curl -s https://api.github.com/repos/sass/node-sass/releases/latest | jq -r '.assets[].browser_download_url')
do
    FILENAME="${URL##*/}"
    if [ $PLATFORM = "all" ]; then
        TARGET=$URL
    else
        TARGET=$(echo $URL | grep $PLATFORM)
    fi

    if [ ! -z $TARGET ]
    then
        if [ -e "$DEST/$FILENAME" ]
        then
            echo "Skipping $FILENAME"
        else
            wget -N $TARGET -P $DEST -q --show-progress
        fi
    fi
done

echo "Setting SASS_BINARY_DIR for terminal session"

export SASS_BINARY_DIR=$DEST

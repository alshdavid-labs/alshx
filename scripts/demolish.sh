#!/bin/bash
if [ -z "$1" ] ; then
    echo "--------"
    echo "Demolish"
    echo "--------"
    echo "Seek and destroy files or folders"
    echo
    echo "Useage: demolish [dir] ...[file/folder]"
    echo "eg: demolish . node_modules .DS_Store thing"
    echo
    exit 1
fi

if [ -z "$2" ] ; then
    echo "--------"
    echo "Demolish"
    echo "--------"
    echo "Seek and destroy files or folders"
    echo
    echo "Useage: demolish [dir] ...[file/folder]"
    echo "eg: demolish . node_modules .DS_Store thing"
    echo
    exit 1
fi

i=0
for var in "$@"
do
    if [ "$i" != 0 ]
    then
        find $1 -name "$var" -exec rm -rf '{}' +
    fi
    i=$(( $i + 1 ))
done
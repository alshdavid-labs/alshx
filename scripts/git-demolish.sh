#!/bin/sh

if [ "$1" = "-h" ] ; then
    echo "--------"
    echo "Git Demolish"
    echo "--------"
    echo "Seek and destroy local git branches"
    echo
    echo "Will ignore master by default"
    echo
    echo "Useage: git-demolish [exclude-braches]"
    echo "eg: git-demolish development"
    echo
    exit 1
fi

INVERT_MATCH="master"
INVERT_MATCH="$INVERT_MATCH\|$1"
i=0
for var in "$@"
do
    if [ "$i" != 0 ]
    then
        INVERT_MATCH="$INVERT_MATCH\|$var"
    fi
    i=$(( $i + 1 ))
done
INVERT_MATCH="$INVERT_MATCH\|*"
git branch | grep --invert-match $INVERT_MATCH | xargs git branch -D

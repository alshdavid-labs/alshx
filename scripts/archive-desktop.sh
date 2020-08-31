#!/bin/bash

if [ "$(ls $HOME/Desktop)" = "" ]; then
  echo "Nothing to copy"
  exit 0
fi

DESKTOP=$HOME

if [ "$ARCHIVE_DESKTOP_SRC" != "" ]; then
  DESKTOP=$ARCHIVE_DESKTOP_SRC
fi

if [ "$ARCHIVE_DESKTOP_DEST" != "" ]; then
  DEST=$ARCHIVE_DESKTOP_DEST
fi

if [ "$1" != "" ]; then
  DEST=$1
fi

if [ -z "$DEST" ]; then
    echo "---------------"
    echo "Archive Desktop"
    echo "---------------"
    echo "Move your desktop to a folder by date"
    echo
    echo "Useage: archive-desktop [dest]"
    echo "eg: archive-desktop ~/Archive"
    echo
    echo "You can also set a variable:"
    echo "ARCHIVE_DESKTOP_DEST=\"\$HOME/Archive\"" 
    exit 1
fi

DAY=`date +%d`
MONTH=`date +%m`
YEAR=`date +%y`
OUT="$DEST/$YEAR.$MONTH.$DAY"

if [ -d "$OUT" ]; then
  mv $DESKTOP/Desktop/* $OUT
else
  mkdir -p $OUT
  mv $DESKTOP/Desktop/* $OUT
fi

echo "Desktop has been moved to: \"$OUT\""

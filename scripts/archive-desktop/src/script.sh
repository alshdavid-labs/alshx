#!/bin/bash

if [ "$ARCHIVE_DESKTOP_SRC" != "" ] && [ -d "$ARCHIVE_DESKTOP_SRC" ]; then
  DESKTOP="$ARCHIVE_DESKTOP_SRC"
elif [ -d "$HOME/Desktop" ]; then
  DESKTOP="$HOME/Desktop"
else
  echo "---------------"
  echo "Archive Desktop"
  echo "---------------"
  echo "Move your desktop to a folder by date"
  echo
  echo "Useage: archive-desktop [dest]"
  echo "eg: archive-desktop ~/Archive"
  echo
  echo  "ERROR"
  echo  "Desktop directory not found"
  echo
  echo "You can set a variables:"
  echo "ARCHIVE_DESKTOP_DEST=\"\$HOME/Archive\"" 
  echo "ARCHIVE_DESKTOP_SRC=\"/mnt/c/Users/alshd/Desktop\""
  exit 1
fi

if [ "$(ls $DESKTOP)" = "" ]; then
  echo "Nothing to copy"
  exit 0
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
  echo  "ERROR"
  echo  "Output directory not found"
  echo
  echo "You can set a variables:"
  echo "ARCHIVE_DESKTOP_DEST=\"\$HOME/Archive\"" 
  echo "ARCHIVE_DESKTOP_SRC=\"/mnt/c/Users/alshd/Desktop\""
  exit 1
fi

DAY=`date +%d`
MONTH=`date +%m`
YEAR=`date +%y`
OUT="$DEST/$YEAR.$MONTH.$DAY"

echo "Moving $DESKTOP to $OUT"

if [ -d "$OUT" ]; then
  mv $DESKTOP/* $OUT
else
  mkdir -p $OUT
  mv $DESKTOP/* $OUT
fi

echo "Desktop has been moved to: \"$OUT\""

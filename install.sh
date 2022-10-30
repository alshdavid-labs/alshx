#!/bin/sh

LINK="https://github.com/alshdavid/alshx/releases/latest/download"
FILENAME="NULL"
OUTDIR="/usr/local/alshx"

OS_TYPE="$(uname -s)"
CPU_TYPE="$(uname -m)"

if [ "$OS_TYPE" = Linux ] && [ "$CPU_TYPE" = "arm64" ]; then
  FILENAME="alshx-linux-arm64.tar.gz"
fi

if [ "$OS_TYPE" = Linux ] && [ "$CPU_TYPE" = "arm" ]; then
  FILENAME="alshx-linux-arm.tar.gz"
fi

if [ "$OS_TYPE" = Linux ] && [ "$CPU_TYPE" = "x86_64" ]; then
  FILENAME="alshx-linux-amd64.tar.gz"
fi

if [ "$OS_TYPE" = Linux ] && [ "$CPU_TYPE" = "x86-64" ]; then
  FILENAME="alshx-linux-amd64.tar.gz"
fi

if [ "$OS_TYPE" = Linux ] && [ "$CPU_TYPE" = "amd64" ]; then
  FILENAME="alshx-linux-amd64.tar.gz"
fi

if [ "$OS_TYPE" = Darwin ] && [ "$CPU_TYPE" = "x86_64" ]; then
  FILENAME="alshx-darwin-amd64.tar.gz"
fi

if [ "$OS_TYPE" = Darwin ] && [ "$CPU_TYPE" = "arm64" ]; then
  FILENAME="alshx-darwin-arm64.tar.gz"
fi

if [ "$FILENAME" = "" ]; then
  echo "Unable to identify computer type"
  exit 1
fi

if [ -d "$OUTDIR" ]; then
  echo "Tools Detected, Upgrading"
  sudo rm -rf "$OUTDIR"
fi

sudo mkdir "$OUTDIR"

if [ -x "$(command -v curl)" ]; then
  curl -s -o "$TMPDIR/$FILENAME"  "$LINK/$FILENAME"
elif [ -x "$(command -v wget)" ]; then
  wget --quiet -O "$TMPDIR/$FILENAME" "$LINK/$FILENAME"
else
  echo "Need either Curl or Wget"
  exit 1
fi

if ! [ "$?" = "0" ]; then
  echo Failed to download file
  exit 1
fi

sudo tar -C "$OUTDIR" -xzf "$TMPDIR/$FILENAME"
if ! [ "$?" = "0" ]; then
  echo Failed to extract file
  exit 1
fi

rm -r -f "$TMPDIR/$FILENAME"

export PATH=/usr/local/alshx:$PATH

echo "add this to your profile: (/etc/profile or /etc/zprofile)"
echo "export PATH=$OUTDIR:\$PATH"

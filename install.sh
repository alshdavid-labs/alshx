#!/bin/sh

echo "Alshx Linux and MacOS Installer"
echo 

echo "Script requires sudo permissions"
sudo echo &> /dev/null

if ! [ "$?" = "0" ]; then
  echo "Failed to obtain sudo permissions"
  exit 1
fi

echo

LINK="https://github.com/alshdavid/alshx/releases/latest/download"
FILENAME="NULL"
OUT_DIR="/usr/local/alshx"
TEMP_DIR="$TMPDIR"

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

if [ "$FILENAME" = "NULL" ]; then
  echo "Unable to identify computer type"
  exit 1
fi

if ! [ "$TMPDIR" = "" ]; then
  TEMP_DIR="$TMPDIR"
elif [ -d "/tmp" ]; then
  TEMP_DIR="/tmp/"
else
  echo "Unable to find temp directory"
  exit 1
fi

echo "Temp directory: \"$TEMP_DIR\""
echo "Downloading archive: \"$FILENAME\""
if [ -x "$(command -v curl)" ]; then
  sudo curl -L -s -o "$TEMP_DIR$FILENAME"  "$LINK/$FILENAME"
elif [ -x "$(command -v wget)" ]; then
  sudo wget --quiet -O "$TEMP_DIR$FILENAME" "$LINK/$FILENAME"
else
  echo "Need either Curl or Wget"
  exit 1
fi

if ! [ "$?" = "0" ]; then
  echo Failed to download file
  exit 1
fi

if [ -d "$OUT_DIR" ]; then
  echo "Alshx already installed"
  echo "Removing directory: \"$OUT_DIR\""
  sudo rm -rf "$OUT_DIR"
fi

echo "Creating directory: \"$OUT_DIR\""
sudo mkdir "$OUT_DIR"

echo "Unpacking archive \"$FILENAME\""
sudo tar -C "$OUT_DIR" -xzf "$TEMP_DIR$FILENAME"
if ! [ "$?" = "0" ]; then
  echo Failed to extract file
  exit 1
fi

echo "Removing archive \"$TEMP_DIR$FILENAME\""
sudo rm -r -f "$TEMP_DIR$FILENAME"

export PATH=/usr/local/alshx:$PATH

echo
echo "Add this to your profile: (/etc/profile or /etc/zprofile)"
echo "    export PATH=$OUT_DIR:\$PATH"
echo

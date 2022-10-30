#!/bin/bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
ROOT_DIR="$( cd "$( dirname "${SCRIPT_DIR[0]}/../../.." )" &> /dev/null && pwd )"
cd "$ROOT_DIR"

rm -rf "${ROOT_DIR}/dist"

RELEASE_DATE=$(date -u --iso-8601=seconds)

env GOOS=linux GOARCH=arm PROD=true make
env GOOS=linux GOARCH=arm64 PROD=true make
env GOOS=linux GOARCH=amd64 PROD=true make
# env GOOS=linux GOARCH=386 PROD=true make

env GOOS=darwin GOARCH=amd64 PROD=true make
env GOOS=darwin GOARCH=arm64 PROD=true make

env GOOS=windows GOARCH=amd64 PROD=true make
# env GOOS=windows GOARCH=386 PROD=true make

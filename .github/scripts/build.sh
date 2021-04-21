#!/bin/bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
ROOT_DIR="$( cd "$( dirname "${SCRIPT_DIR[0]}/../../.." )" &> /dev/null && pwd )"

YEAR="$(date -u +%y)"
MONTH="$(date -u +%m)"
DAY="$(date -u +%d)"
HOUR="$(date -u +%H)"
MINUTE="$(date -u +%M)"
VERSION="${YEAR}-${MONTH}-${DAY}-${HOUR}-${MINUTE}"
echo "Tag: ${VERSION}"

rm -rf "${ROOT_DIR}/dist"
cd "${ROOT_DIR}/tools"

echo "Building Packages..."

for d in */
do
  cd "${ROOT_DIR}/tools/${d}"

  NAME=$(basename "${PWD}")
  OUT="${ROOT_DIR}/dist"

  echo "Building package at ${PWD}"
  env GOOS=linux GOARCH=arm go build -ldflags="-s -w" -tags prod -o "${OUT}/linux-arm32/${NAME}" ./src
  env GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -tags prod -o "${OUT}/linux-arm64/${NAME}" ./src
  env GOOS=linux GOARCH=386 go build -ldflags="-s -w" -tags prod -o "${OUT}/linux-386/${NAME}" ./src
  env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -tags prod -o "${OUT}/linux-amd64/${NAME}" ./src
  env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -tags prod -o "${OUT}/darwin-amd64/${NAME}" ./src
  env GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -tags prod -o "${OUT}/darwin-arm64/${NAME}" ./src
  env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -tags prod -o "${OUT}/windows-amd64/${NAME}.exe" ./src
  env GOOS=windows GOARCH=386 go build -ldflags="-s -w" -tags prod -o "${OUT}/windows-386/${NAME}.exe" ./src
done

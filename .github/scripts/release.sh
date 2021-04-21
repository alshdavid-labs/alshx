#!/bin/sh
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
ROOT_DIR="$( cd "$( dirname "${SCRIPT_DIR[0]}/../../.." )" &> /dev/null && pwd )"

cd "${ROOT_DIR}/dist"

for i in */
do
  cd "${ROOT_DIR}/dist/$i"
  NAME=$(basename "${PWD}")

  zip -r -D "${ROOT_DIR}/dist/${NAME}.zip" *
  rm -rf "${ROOT_DIR}/dist/${i}"
done

cd "${ROOT_DIR}"

gh release create "${VERSION}" -n "Automatically generated binaries" ./dist/*.zip
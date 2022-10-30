#!/bin/bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
ROOT_DIR="$( cd "$( dirname "${SCRIPT_DIR[0]}/../../.." )" &> /dev/null && pwd )"

TAG="$(date -u +"v%Y.%m.%d.%H%M")"

echo "Tag: ${TAG}"

cd "${ROOT_DIR}/dist"

echo "Publishing Release..."
for i in */
do
  cd "${ROOT_DIR}/dist/$i"
  NAME=$(basename "${PWD}")

  tar -czvf "${ROOT_DIR}/dist/alshx-${NAME}.tar.gz" *
  rm -rf "${ROOT_DIR}/dist/${i}"
done

cd "${ROOT_DIR}"

if [ "$CI" = "true" ]; then
  gh release create "${TAG}" -n "Automatically generated binaries" ./dist/*.tar.gz
fi

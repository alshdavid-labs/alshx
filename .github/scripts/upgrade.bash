#!/bin/bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
ROOT_DIR="$( cd "$( dirname "${SCRIPT_DIR[0]}/../../.." )" &> /dev/null && pwd )"

echo "Tag: ${VERSION}"

rm -rf "${ROOT_DIR}/dist"
cd "${ROOT_DIR}/tools"

echo "Building Packages..."

if [ "$CI" = "true" ]; then
  find . -type f -exec sed -i "s/__VERSION__/${VERSION}/" {} +
fi

for d in */
do
  cd "${ROOT_DIR}/tools/${d}"

  NAME=$(basename "${PWD}")
  OUT="${ROOT_DIR}/dist"

  echo "Updating package dependencies at ${PWD}"

  go get -u all
done

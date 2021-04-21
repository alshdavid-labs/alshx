#!/bin/sh
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
ROOT_DIR="$( cd "$( dirname "${SCRIPT_DIR[0]}/../../.." )" &> /dev/null && pwd )"

cd "${ROOT_DIR}/tools"

for d in */
do
  cd "${ROOT_DIR}/tools/${d}"
  echo "TODO: Testing ${PWD}"
done

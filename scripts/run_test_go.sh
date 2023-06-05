#!/usr/bin/env bash

set -e;

current_path=$(cd "$(dirname "${0}")" || exit 1; pwd)
cd "${current_path}"/.. || exit 1

CASE="$1"
DIR="$2"

ROOT="./xgo/tests/cases"

function runCases() {
  dir="$1"

  go test -v -test.count=1 -failfast -test.run="${CASE}" "${dir}/..."
  return
}

if [ -n "${DIR}" ]; then
  runCases "${ROOT}/${DIR}"
  exit $?
fi

# run test cases in boundary.
runCases "${ROOT}/boundary"

# run test cases in types.
runCases "${ROOT}/types"

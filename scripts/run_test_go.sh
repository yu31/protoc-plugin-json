#!/usr/bin/env bash

set -e;

current_path=$(cd "$(dirname "${0}")" || exit 1; pwd)
cd "${current_path}"/.. || exit 1

BASE_DIR="./xgo/tests/cases"

# LEVEL0/CASE is set by environment variables;
# e.g: LEVEL0=advanced CASE=Test_WKT_Any1 make test
LEVEL0="${LEVEL0}"
CASE="${CASE}"

if [ -z "${LEVEL0}" ]; then
  LEVEL0="*"
fi

function runCases() {
  test_dir="$1"

  go test -v -test.count=1 -failfast -test.run="${CASE}" "${test_dir}/..."
  return
}

for dir in "${BASE_DIR}"/${LEVEL0}; do
  runCases "${dir}"
done

go test -v -test.count=1 -test.failfast -test.run="${CASE}" ./xgo/pkg/jsonencoder/
go test -v -test.count=1 -test.failfast -test.run="${CASE}" ./xgo/pkg/jsondecoder/

#!/usr/bin/env bash


CURRENT_PATH=$(cd "$(dirname "${0}")" || exit 1; pwd)
cd "${CURRENT_PATH}"/.. || exit 1

export PATH="./build/bin:$PATH"

PLUGIN="gojson"
BASE_DIR="tests/proto/invalid"

# DEBUG/LEVEL0/LEVEL1/FILE is set by environment variables;
# e.g: DEBUG=yes LEVEL0=advanced LEVEL1=empty FILE=empty_key1.proto make test-invalid
DEBUG="${DEBUG}"
LEVEL0="${LEVEL0}"
LEVEL1="${LEVEL1}"
FILE="${FILE}"

if [ -z "${LEVEL0}" ]; then
  LEVEL0="*"
fi
if [ -z "${LEVEL1}" ]; then
  LEVEL1="*"
fi
if [ -z "${FILE}" ]; then
  FILE="*"
fi

function runCase() {
  file="$1"
  echo "=== RUN: ${file}"

  ret=$(protoc -I=. --go_out=./ --"${PLUGIN}"_out=./ "${file}" 2>&1)

  if [ "${DEBUG}" == "yes" ]; then
    echo "${ret}"
  fi

  if echo "${ret}" |grep "\[ERROR\] - \[plugin: ${PLUGIN}\]" > /dev/null; then
    echo "--- PASS: ${file}"
  else
    echo "    Expected error by ${PLUGIN}, but it passing or other error occurs"
    echo "--- FAIL: ${file}"
    exit 1
  fi
  return
}

for file in ${BASE_DIR}/${LEVEL0}/${LEVEL1}/${FILE}; do
  runCase "${file}"
done

echo "--------------------------------------------"
echo "------ ALL test cases are successful in language golang ------"
echo "--------------------------------------------"

# Delete generated files.
/bin/rm -fr ./xgo/tests/pb/pbinvalid/*;

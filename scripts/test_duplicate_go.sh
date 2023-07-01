#!/usr/bin/env bash


CURRENT_PATH=$(cd "$(dirname "${0}")" || exit 1; pwd)
cd "${CURRENT_PATH}"/.. || exit 1

export PATH="./build/bin:$PATH"

PLUGIN="gojson"
BASE_DIR="tests/proto/duplicate"

FILE="*${1}*"

#if [ -z "${FILE}" ]; then
#  FILE="*"
#fi

function runCase() {
  echo "=== RUN:  ${file}"

  ret=$(protoc -I=. --go_out=./ --"${PLUGIN}"_out=./ "${file}" 2>&1)

  if [ "${DEBUG}" != "no" ]; then
    echo "${ret}"
  fi

#  if echo "${ret}" |grep "ERROR - gojson" > /dev/null; then
#    echo "--- PASS: ${file}"
#  else
#    echo "    Expected error by godefault, but it passing or other error occurs"
#    echo "--- FAIL: ${file}"
#    exit 1
#  fi
  return
}

for file in ${BASE_DIR}/${FILE}; do
  runCase "${file}"
done

# Delete generated files.
/bin/rm -fr ./xgo/tests/pb/pbduplicate/*;

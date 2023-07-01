#!/usr/bin/env bash


CURRENT_PATH=$(cd "$(dirname "${0}")" || exit 1; pwd)
cd "${CURRENT_PATH}"/.. || exit 1

export PATH="./build/bin:$PATH"

PLUGIN="gojson"
BASE_DIR="tests/proto/invalid"

KIND="$1"

if [ -z "${KIND}" ]; then
  KIND="*"
fi

function runCase() {
  echo "=== RUN:  ${file}"

  ret=$(protoc -I=. --go_out=./ --"${PLUGIN}"_out=./ "${file}" 2>&1)

  if [ "${DEBUG}" != "no" ]; then
    echo "${ret}"
  fi

  if echo "${ret}" |grep '\[ERROR\] - \[plugin: gojson\]' > /dev/null; then
    echo "--- PASS: ${file}"
  else
    echo "    Expected error by gojson, but it passing or other error occurs"
    echo "--- FAIL: ${file}"
    exit 1
  fi
  return
}

for file in ${BASE_DIR}/${KIND}/*; do
  runCase "${file}"
done

echo "--------------------------------------------"
echo "------ ALL test cases are successful ------"
echo "--------------------------------------------"

# Delete generated files.
/bin/rm -fr ./xgo/tests/pb/pbinvalid/*;

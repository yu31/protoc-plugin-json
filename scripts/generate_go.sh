#!/usr/bin/env bash

set -e;

CURRENT_PATH=$(cd "$(dirname "${0}")" || exit 1; pwd)

# install dep package.
sh "${CURRENT_PATH}/ensure_dep.sh"

cd "${CURRENT_PATH}"/.. || exit 1

output_dir="./xgo/pb"
# To avoids invalid code residue.
/bin/rm -fr "$output_dir"
mkdir -p "$output_dir"

MODULE="$(go list)"

for f in proto/*.proto; do
#  echo "Generated Golang code for proto file ${f}"
  protoc -I=./proto --go_opt=module="${MODULE}" --go_out=. "$f"
done

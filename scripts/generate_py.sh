#!/usr/bin/env bash
# Generate protobuf code for python

set -e;

if ! [[ "$0" =~ scripts/generate_py.sh ]]; then
	echo "must be run from repository root"
	exit 255
fi

CURRENT_PATH=$(cd "$(dirname "${0}")" || exit 1; pwd)

# install dep package.
sh "${CURRENT_PATH}/ensure_dep.sh"

cd "${CURRENT_PATH}"/.. || exit 1

GOPATH=$(go env GOPATH)

output_dir="./xpy/src/pb"

# To avoids invalid code residue.
/bin/rm -fr "$output_dir"
mkdir -p "$output_dir"

for f in proto/*.proto; do
    # generate python code.
#    echo "Generated Python code for proto file ${f}"
  protoc -I. -I./proto -I"${GOPATH}"/src --python_out="$output_dir" "${f}"
done


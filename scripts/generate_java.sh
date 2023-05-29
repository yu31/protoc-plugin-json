#!/usr/bin/env bash

set -e;

CURRENT_PATH=$(cd "$(dirname "${0}")" || exit 1; pwd)

# install dep package.
sh "${CURRENT_PATH}/ensure_dep.sh"

cd "${CURRENT_PATH}"/.. || exit 1

output_dir="./xjava/src/main/java"

# To avoids invalid code residue.
/bin/rm -fr "${output_dir}/io/github/yu31/protoc/pb"
mkdir -p "$output_dir"

for f in proto/*.proto; do
  # code for java
#  echo "Generated Java code for proto file ${f}"
  protoc -I./proto --java_out="$output_dir" "$f"
done

#if git status |grep 'src/main/java' >/dev/null; then
#  echo "mvn clean package deploy"
#  mvn clean package deploy >/dev/null 2>&1 || exit $?
#else
#  echo "no java code generated, skip mvn deploy"
#fi

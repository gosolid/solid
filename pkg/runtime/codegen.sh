#!/bin/bash

PROJECT_ROOT="$(cd -P -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)"

for FOLDER in ${PROJECT_ROOT}/../task ${PROJECT_ROOT}/crypto/web ${PROJECT_ROOT}/*; do
  if [ -d "${FOLDER}" ]; then
    go generate github.com/grexie/isolates/codegen "${FOLDER}"
    echo "✅ $(basename "${FOLDER}")"
  fi
done
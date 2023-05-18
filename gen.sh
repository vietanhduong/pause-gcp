#!/bin/bash
set -eu

paths=(
  apis
)

_protos=""
for p in "${paths[@]}"; do
  _protos+="--path ${p} "
done

buf format -w $_protos
buf lint $_protos || exit 1

# Generate all protos
buf generate $_protos

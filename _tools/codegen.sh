#!/bin/bash

DIR=$(cd $(dirname $0); pwd)
KUBE_DIR=$(cd $(dirname $(dirname $0)); pwd)/kube

# clean generated files
rm ${KUBE_DIR}/*.gen.go

set -e

subcmds=(
    "get"
    "describe"
    "create"
    "replace"
    "patch"
    "delete"
    "edit"
    "apply"
    "logs"
    "scale"
    "attach"
    "exec"
    "proxy"
    "run"
    "expose"
    "label"
    "explain"
    "cordon"
    "drain"
    "uncordon"
    "annotate"
    "convert"
)

for cmd in "${subcmds[@]}"; do
  kubectl ${cmd} --help | go run ${DIR}/resource-gen/main.go -o ${KUBE_DIR}/option_${cmd}.gen.go -var ${cmd}Options
  goimports -w ${KUBE_DIR}/option_${cmd}.gen.go
done

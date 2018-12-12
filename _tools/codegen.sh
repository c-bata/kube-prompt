#!/bin/bash

DIR=$(cd $(dirname $0); pwd)
KUBE_DIR=$(cd $(dirname $(dirname $0)); pwd)/kube

# clean generated files
rm ${KUBE_DIR}/*.gen.go
mkdir -p bin

set -e

go build -o ./bin/option-gen ./_tools/option-gen/main.go

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
    "rolling-update"
    "scale"
    "attach"
    "exec"
    "port-forward"
    "proxy"
    "run"
    "expose"
    "autoscale"
    "label"
    "explain"
    "cordon"
    "drain"
    "uncordon"
    "annotate"
    "convert"
    "cluster-info"
)

for cmd in "${subcmds[@]}"; do
  camelized=`echo ${cmd} | gsed -r 's/-(.)/\U\1\E/g'`
  kubectl ${cmd} --help | ./bin/option-gen -o ${KUBE_DIR}/option_${cmd//-/_}.gen.go -var ${camelized}Options
  goimports -w ${KUBE_DIR}/option_${cmd//-/_}.gen.go
done

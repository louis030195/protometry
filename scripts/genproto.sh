#!/usr/bin/env bash
#
# https://github.com/prometheus/prometheus/blob/2bd510a63e48ac6bf4971d62199bdb1045c93f1a/scripts/genproto.sh
# Generate all protobuf bindings.
# Run from repository root.
set -e
set -u

if ! [[ "$0" =~ "scripts/genproto.sh" ]]; then
	echo "must be run from repository root"
	exit 255
fi

if ! [[ $(protoc --version) =~ "3.7" ]]; then
	echo "could not find protoc 3.7.x, is it installed + in PATH?"
	exit 255
fi

PROT_ROOT="${GOPATH}/src/github.com/louis030195/protometry"
PROT_PATH="${PROT_ROOT}/api"

DIRS=("api")

for dir in ${DIRS[*]}; do
	pushd "${dir}"
		protoc --go_out=../pkg -I=. \
            -I="${PROT_PATH}" \
            *.proto

    # Hack to use the lib, not the proto, required for this kind of stuff: (box.Min.Max(vector.Vector3{})
#    sed -i 's#vector "vector"#"github.com/louis030195/protometry/pkg/vector"#g' ../../pkg/"${dir}"/*.pb.go
		#goimports -w *.pb.go
	popd
done

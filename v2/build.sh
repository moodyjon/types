#!/bin/bash

set -euo pipefail
#set -x

version_gte() {
  [ "$1" = "$(echo -e "$1\n$2" | sort -V | tail -n1)" ]
}

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"


hash protoc 2>/dev/null || { echo >&2 -e 'error: protoc binary not found\nDownload it from https://github.com/google/protobuf/releases and put it in your path.\nMake sure you get the one starting with `protoc`, not `protobuf`.'; exit 1; }


PROTOC="$(which protoc)"
VERSION="$($PROTOC --version | cut -d' ' -f2)"
MIN_VERSION="3.0"

version_gte "$VERSION" "$MIN_VERSION" || { echo >&2 "error: protoc version must be >= $MIN_VERSION (your $PROTOC is $VERSION)"; exit 1; }


hash protoc-gen-go 2>/dev/null || go get -u github.com/golang/protobuf/protoc-gen-go
hash protoc-gen-go 2>/dev/null || { echo >&2 'error: Make sure $GOPATH/bin is in your $PATH'; exit 1; }

# Inform protoc what package Go code should be placed within.
GO_PACKAGE="github.com/lbryio/herald.go/protobuf/go"
GO_OPTS+=" --go_opt=Mclaim.proto=$GO_PACKAGE"
GO_OPTS+=" --go_opt=Mpurchase.proto=$GO_PACKAGE"
GO_OPTS+=" --go_opt=Mresult.proto=$GO_PACKAGE"
GO_OPTS+=" --go_opt=Msupport.proto=$GO_PACKAGE"
GO_OPTS+=" --go_opt=Mstringmap_ext.proto=$GO_PACKAGE"

mkdir -p $DIR/go $DIR/go.tmp $DIR/python $DIR/js $DIR/cpp
find $DIR/go $DIR/go.tmp $DIR/python $DIR/js $DIR/cpp -type f -delete


protoc --proto_path="$DIR/proto" \
    --python_out="$DIR/python" \
    --go_out="$DIR/go.tmp" $GO_OPTS \
    --js_out="import_style=commonjs,binary:$DIR/js" \
    --cpp_out="$DIR/cpp" \
    $DIR/proto/*.proto

# Fixup generated Go code.
ls "$DIR"/go.tmp/$GO_PACKAGE/*.pb.go | xargs -n1 -IX -S1024 bash -c "sed -e 's/,omitempty//' X > X.tmp && mv X{.tmp,}"
cp "$DIR"/go.tmp/$GO_PACKAGE/* "$DIR"/go/
rm -R "$DIR"/go.tmp

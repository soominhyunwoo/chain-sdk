#!/usr/bin/env bash

set -eo pipefail

protoc_gen_gosoominhyunwoo() {
  if ! grep "github.com/gogo/protobuf => github.com/regen-network/protobuf" go.mod &>/dev/null ; then
    echo -e "\tPlease run this command from somewhere inside the soominhyunwoo-sdk folder."
    return 1
  fi

  go get github.com/soominhyunwoo/chain-proto/protoc-gen-gosoominhyunwoo@latest 2>/dev/null
}

protoc_gen_gosoominhyunwoo

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  buf protoc \
    -I "proto" \
    -I "third_party/proto" \
    --gosoominhyunwoo_out=plugins=interfacetype+grpc,\
Mgoogle/protobuf/any.proto=github.com/soominhyunwoo/chain-sdk/codec/types:. \
    --grpc-gateway_out=logtostderr=true,allow_colon_final_segments=true:. \
  $(find "${dir}" -maxdepth 1 -name '*.proto')

done

# command to generate docs using protoc-gen-doc
buf protoc \
  -I "proto" \
  -I "third_party/proto" \
  --doc_out=./docs/core \
  --doc_opt=./docs/protodoc-markdown.tmpl,proto-docs.md \
  $(find "$(pwd)/proto" -maxdepth 5 -name '*.proto')
go mod tidy

# generate codec/testdata proto code
buf protoc -I "proto" -I "third_party/proto" -I "testutil/testdata" --gosoominhyunwoo_out=plugins=interfacetype+grpc,\
Mgoogle/protobuf/any.proto=github.com/soominhyunwoo/chain-sdk/codec/types:. ./testutil/testdata/*.proto

# move proto files to the right places
cp -r github.com/soominhyunwoo/chain-sdk/* ./
rm -rf github.com

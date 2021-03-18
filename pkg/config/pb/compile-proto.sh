#!/bin/bash

# Get the full directory name of the script no matter where it is being called from
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Root of mono-repo
ROOT_DIR=$DIR/../../../

PROTO_DIR=$ROOT_DIR/proto
PB_OUT_DIR=$DIR

# Move to $PROTO_DIR
cd $PROTO_DIR

# Compile .proto files
protoc --go_out=$PB_OUT_DIR ./config/v1/*.proto
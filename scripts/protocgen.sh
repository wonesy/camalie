#!/bin/bash

SVCS=("chat")

for svc in "${SVCS[@]}"; do
    echo "Processing service ${svc}"
    protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./proto/${svc}/*.proto
done
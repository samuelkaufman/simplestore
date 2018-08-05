#!/bin/bash
set -x
set errexit
set nounset
export GOPATH=$(mktemp -d)
cp -r vendor $GOPATH/src
go get github.com/samuelkaufman/simplestore/pkg/simplestoreapp
cd cmd/simplestore
gcloud app deploy

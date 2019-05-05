#!/bin/bash
git clone https://github.com/golang/go/
GODEV=$(pwd)/go
cd go/src
GOROOT_BOOTSTRAP=$(go env GOROOT) ./make.bash
$GODEV/bin/go test -bench=. -v sort -run TestSearch

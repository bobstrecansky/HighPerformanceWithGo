#!/bin/bash

for i in {1..4..1}
do
    export GOMAXPROCS=$i
    printf "\nBuild with GOMAXPROCS=$i:"
    time go build -a std
done

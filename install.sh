#!/bin/bash

if [ -d $GOPATH ]; then
    cp -r GameOfLife $GOPATH/src/
else
    echo "Please set "'$GOPATH'" - https://github.com/golang/go/wiki/Setting-GOPATH"
fi

go install GameOfLife
# for tomorrow: error with info about setting up $GOSRCPATH if above fails

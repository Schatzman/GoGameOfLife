#!/bin/bash

if [ -d $GOSRCPATH ]; then
    cp -r GameOfLife $GOSRCPATH/
else
    echo "Please set "'$GOSRCPATH' should be `$GOPATH/src`" - https://github.com/golang/go/wiki/Setting-GOPATH"
fi

go install GameOfLife

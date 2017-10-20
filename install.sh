#!/bin/bash

if [ -d $GOSRCPATH ]; then
    cp -r GameOfLife $GOSRCPATH/src/
else
    echo "Please set "'$GOSRCPATH' should be `$GOPATH/src`" - https://github.com/golang/go/wiki/Setting-GOPATH"
fi

go install GameOfLife
# for tomorrow: error with info about setting up $GOSRCPATH if above fails

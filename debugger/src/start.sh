#!/bin/sh

go install --gcflags="-N -l" github.com/jazzyfresh/hello
go install github.com/jazzyfresh/debugger

#!/bin/sh

go env -w GO111MODULE=off
go install --gcflags="-N -l" github.com/jazzyfresh/hello
go install github.com/jazzyfresh/debugger

# chmod +x start.sh

#!/bin/bash

rm -rf ./build

mkdir build

cd ./pkg/cli
GOOS=darwin GOARCH=386 go build -o ../../build/glow_darwin_386 .
GOOS=darwin GOARCH=amd64 go build -o ../../build/glow_darwin_amd64 .
GOOS=linux GOARCH=386 go build -o ../../build/glow_linux_386 .
GOOS=linux GOARCH=amd64 go build -o ../../build/glow_linux_amd64 .
GOOS=windows GOARCH=386 go build -o ../../build/glow_windows_386.exe .
GOOS=windows GOARCH=amd64 go build -o ../../build/glow_windows_amd64.exe .

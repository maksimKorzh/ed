#!/bin/bash
export GOOS=windows && go build -o edit.exe *.go
export GOOS=linux && go build -o edit *.go

#!/bin/bash
# Chapter 3: Go Module Init & Dependency Management
mkdir mymodule && cd mymodule
go mod init github.com/yourname/mymodule
go get github.com/pkg/errors@v0.9.1
go mod tidy
go mod vendor

#!/bin/sh

heading() {
    echo
    echo "\x1B[32m\x1B[1m$1 (Installing)\x1B[0m"
}

heading go-outline
go get -u -v github.com/ramya-rao-a/go-outline

heading go-symbols
go get -u -v github.com/acroca/go-symbols

heading gocode
go get -u -v github.com/mdempsky/gocode

heading godef
go get -u -v github.com/rogpeppe/godef

heading godoc
go get -u -v golang.org/x/tools/cmd/godoc

heading gogetdoc
go get -u -v github.com/zmb3/gogetdoc

heading golint
go get -u -v golang.org/x/lint/golint

heading gomodifytags
go get -u -v github.com/fatih/gomodifytags

heading gorename
go get -u -v golang.org/x/tools/cmd/gorename

heading goreturns
go get -u -v sourcegraph.com/sqs/goreturns

heading goimports
go get -u -v golang.org/x/tools/cmd/goimports

heading gotests
go get -u -v github.com/cweill/gotests/...

heading guru
go get -u -v golang.org/x/tools/cmd/guru

heading impl
go get -u -v github.com/josharian/impl

heading goplay
go get -u -v github.com/haya14busa/goplay/cmd/goplay

heading gopkgs
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs

heading fillstruct
go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct

heading gometalinter
go get -u -v github.com/alecthomas/gometalinter

heading gometalinter install
gometalinter --install

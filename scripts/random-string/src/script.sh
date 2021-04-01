#!/bin/sh

LENGTH=$1

if [ -z $1 ]; then
    LENGTH=10
fi

openssl rand -base64 $LENGTH

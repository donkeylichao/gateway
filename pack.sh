#!/bin/sh

tarfile="gateway-$1.tar.gz"

echo "开始打包$tarfile..."

export GOARCH=amd64
export GOOS=linux
#export GOOS=darwin

bee pack

mv gateway.tar.gz $tarfile

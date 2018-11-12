#!/bin/bash

TAG=v0.1.0

git tag -a $TAG -m "$TAG release"
git push origin $TAG

mkdir binaries/releases/$TAG

# mac
env GOOS=darwin GOARCH=amd64 go build -o binaries/mac/eksuser
zip binaries/releases/$TAG/eksuser-darwin-amd64.zip binaries/mac/eksuser
# linux
env GOOS=linux GOARCH=amd64 go build -o binaries/linux/eksuser
zip binaries/releases/$TAG/eksuser-linux-amd64.zip binaries/linux/eksuser

# windows
env GOOS=windows GOARCH=amd64 go build -o binaries/windows/eksuser.exe
zip binaries/releases/$TAG/eksuser-windows-amd64.zip binaries/windows/eksuser.exe

# copy to current path
cp binaries/mac/eksuser ~/bin/

# create a release
hub release create \
    -a binaries/releases/$TAG/eksuser-mac-amd64.zip \
    -a binaries/releases/$TAG/eksuser-linux-amd64.zip \
    -a binaries/releases/$TAG/eksuser-windows-amd64.zip \
    -m "$TAG" $TAG
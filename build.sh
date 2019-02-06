#!/bin/bash

TAG=v0.1.1

git tag -a $TAG -m "$TAG release"
git push origin $TAG

mkdir binaries/releases/$TAG

# darwin
env GOOS=darwin GOARCH=amd64 go build -o binaries/darwin/eksuser
zip binaries/releases/$TAG/eksuser-darwin-amd64.zip binaries/darwin/eksuser

# linux
env GOOS=linux GOARCH=amd64 go build -o binaries/linux/eksuser
zip binaries/releases/$TAG/eksuser-linux-amd64.zip binaries/linux/eksuser

# windows
env GOOS=windows GOARCH=amd64 go build -o binaries/windows/eksuser.exe
zip binaries/releases/$TAG/eksuser-windows-amd64.zip binaries/windows/eksuser.exe

# copy to current path
# cp binaries/darwin/eksuser ~/bin/

# create a release
hub release create \
    -a binaries/releases/$TAG/eksuser-darwin-amd64.zip \
    -a binaries/releases/$TAG/eksuser-linux-amd64.zip \
    -a binaries/releases/$TAG/eksuser-windows-amd64.zip \
    -m "$TAG" $TAG
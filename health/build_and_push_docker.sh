#!/bin/sh

export VERSION=2

./build.sh

docker build -t ragnaroek/health:"$VERSION" .

docker push ragnaroek/health:"$VERSION"

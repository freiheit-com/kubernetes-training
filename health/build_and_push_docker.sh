#!/bin/sh

export REG_NAME="fdc-test-statistic" # you have to change this, if you create a new Kubernetes cluster for the training
export VERSION=2

./build.sh

docker build -t ragnaroek/health:"$VERSION" .

docker push ragnaroek/health:"$VERSION"

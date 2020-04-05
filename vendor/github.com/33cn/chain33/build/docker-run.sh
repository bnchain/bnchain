#!/usr/bin/env bash
# first you must build docker image, you can use make docker command
# docker build . -f Dockerfile-run -t bnchain-build:latest

sudo docker run -it -p 8801:8801 -p 8802:8802 -p 6060:6060 -p 50051:50051 -l linux-bnchain-run \
    -v "$GOPATH"/src/gitlab.bitnasdaqchain.com/bnchain/bnchain:/go/src/gitlab.bitnasdaqchain.com/bnchain/bnchain \
    -w /go/src/gitlab.bitnasdaqchain.com/bnchain/bnchain bnchain:latest

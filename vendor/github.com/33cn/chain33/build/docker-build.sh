#!/usr/bin/env bash
# https://hub.docker.com/r/suyanlong/golang-dev/
# https://github.com/suyanlong/golang-dev
# sudo docker pull suyanlong/golang-dev:latest

sudo docker run -it -p 8801:8801 -p 8802:8802 -p 6060:6060 -p 50051:50051 -l linux-bnchain-build \
    -v "$GOPATH"/src/gitlab.bitnasdaqchain.com/bnchain/bnchain:/go/src/gitlab.bitnasdaqchain.com/bnchain/bnchain \
    -w /go/src/gitlab.bitnasdaqchain.com/bnchain/bnchain suyanlong/golang-dev:latest

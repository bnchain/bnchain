#!/bin/bash

while :; do
    ./bnchain-cli net time
    #nc -vz localhost 8805
    sleep 1
done

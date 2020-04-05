#!/usr/bin/env bash
/root/bnchain -f /root/bnchain.toml &
# to wait nginx start
sleep 15
/root/bnchain -f "$PARAFILE"

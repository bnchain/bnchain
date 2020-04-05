FROM ubuntu:16.04

WORKDIR /root
COPY bnchain bnchain
COPY bnchain-cli bnchain-cli
COPY bnchain.toml ./

CMD ["/root/bnchain", "-f", "/root/bnchain.toml"]

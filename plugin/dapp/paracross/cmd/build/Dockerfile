FROM ubuntu:16.04

WORKDIR /root
COPY bnchain bnchain
COPY bnchain-cli bnchain-cli
COPY bnchain-para-cli bnchain-para-cli
COPY bnchain.toml bnchain*.toml ./
COPY entrypoint.sh entrypoint.sh

CMD ["/root/bnchain", "-f", "/root/bnchain.toml"]

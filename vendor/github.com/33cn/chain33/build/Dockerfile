FROM ubuntu:16.04

WORKDIR /root
COPY bnchain bnchain
COPY bnchain-cli bnchain-cli
COPY bnchain.toml ./

RUN ./bnchain-cli cert --host=127.0.0.1

CMD ["/root/bnchain", "-f", "/root/bnchain.toml"]

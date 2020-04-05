[![pipeline status](https://api.travis-ci.org/bnchain/plugin.svg?branch=master)](https://travis-ci.org/bnchain/plugin/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bnchain/plugin?branch=master)](https://goreportcard.com/report/github.com/bnchain/plugin)


# BitNasdaq chain system based on bnchain blockchain development framework


### Compile

```
git clone https://github.com/bnchain/bnchain $GOPATH/src/github.com/bnchain/bnchain
cd $GOPATH/src/github.com/bnchain/bnchain
go build -i -o bnchain
go build -i -o bnchain-cli github.com/bnchain/bnchain/cli
```

### Function
Copy the three compiled files bnchain, bnchain-cli, bn.toml and put them in the same folder. Execute:
```
./bnchain -f bn.toml
```

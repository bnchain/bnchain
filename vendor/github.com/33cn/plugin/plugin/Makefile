all:
	go get -v -u gitlab.bitnasdaqchain.com/bnchain/bnchain
	go build -i -o tool github.com/bnchain/bnchain/cmd/tools
	./tool import --path "." --packname "github.com/bnchain/plugin" --conf "" --out "plugin.toml"


# golang1.9 or latest
# 1. make help
# 2. make dep
# 3. make build
# ...

CLI := build/bnchain-cli
SRC_CLI := github.com/bnchain/plugin/cli
APP := build/bnchain
bnchain=github.com/bnchain/bnchain
bnchain_PATH=vendor/${bnchain}
LDFLAGS := -ldflags "-w -s"
PKG_LIST_VET := `go list ./... | grep -v "vendor" | grep -v plugin/dapp/evm/executor/vm/common/crypto/bn256`
PKG_LIST := `go list ./... | grep -v "vendor" | grep -v "bnchain/test" | grep -v "mocks" | grep -v "pbft"`
PKG_LIST_INEFFASSIGN= `go list -f {{.Dir}} ./... | grep -v "vendor"`
BUILD_FLAGS = -ldflags "-X github.com/bnchain/plugin/vendor/github.com/bnchain/bnchain/common/version.GitCommit=`git rev-parse --short=8 HEAD`"
MKPATH=$(abspath $(lastword $(MAKEFILE_LIST)))
MKDIR=$(dir $(MKPATH))
proj := "build"
.PHONY: default dep all build release cli linter race test fmt vet bench msan coverage coverhtml docker docker-compose protobuf clean help autotest

default: build depends

build:
	go build $(BUILD_FLAGS) -v -i -o $(APP)
	go build $(BUILD_FLAGS) -v -i -o $(CLI) $(SRC_CLI)
	@cp bnchain.toml $(bnchain_PATH)/build/system-test-rpc.sh build/

build_ci: depends ## Build the binary file for CI
	@go build -v -i -o $(CLI) $(SRC_CLI)
	@go build $(BUILD_FLAGS) -v -o $(APP)
	@cp bnchain.toml $(bnchain_PATH)/build/system-test-rpc.sh build/

para:
	@go build -v -o build/$(NAME) -ldflags "-X $(SRC_CLI)/buildflags.ParaName=user.p.$(NAME). -X $(SRC_CLI)/buildflags.RPCAddr=http://localhost:8901" $(SRC_CLI)

vet:
	@go vet ${PKG_LIST_VET}

autotest: ## build autotest binary
	@cd build/autotest && bash ./build.sh && cd ../../
	@if [ -n "$(dapp)" ]; then \
		rm -rf build/autotest/local \
		&& cp -r $(bnchain_PATH)/build/autotest/local $(bnchain_PATH)/build/autotest/*.sh build/autotest/ \
		&& cd build/autotest && bash ./copy-autotest.sh local && cd local && bash ./local-autotest.sh $(dapp) && cd ../../../; fi
autotest_ci: autotest ## autotest ci
	@rm -rf build/autotest/jerkinsci \
	&& cp -r $(bnchain_PATH)/build/autotest/jerkinsci $(bnchain_PATH)/build/autotest/*.sh build/autotest/ \
	&& cd build/autotest && bash ./copy-autotest.sh jerkinsci/temp$(proj) \
	&& cd jerkinsci && bash ./jerkins-ci-autotest.sh $(proj) && cd ../../../
autotest_tick: autotest ## run with ticket mining
	@rm -rf build/autotest/gitlabci \
	&& cp -r $(bnchain_PATH)/build/autotest/gitlabci $(bnchain_PATH)/build/autotest/*.sh build/autotest/ \
	&& cd build/autotest && bash ./copy-autotest.sh gitlabci \
	&& cd gitlabci && bash ./gitlab-ci-autotest.sh build && cd ../../../

update:
	rm -rf ${bnchain_PATH}
	git clone --depth 1 -b ${b} https://${bnchain}.git ${bnchain_PATH}
	rm -rf vendor/${bnchain}/.git
	rm -rf vendor/${bnchain}/vendor/github.com/apache/thrift/tutorial/erl/
	cp -Rf vendor/${bnchain}/vendor/* vendor/
	rm -rf vendor/${bnchain}/vendor
	govendor init
	go build -i -o tool github.com/bnchain/plugin/vendor/github.com/bnchain/bnchain/cmd/tools
	./tool import --path "plugin" --packname "github.com/bnchain/plugin/plugin" --conf ""

updatevendor:
	govendor add +e
	govendor fetch -v +m

dep:
	dep init -v

linter: vet ineffassign ## Use gometalinter check code, ignore some unserious warning
	@./golinter.sh "filter"
	@find . -name '*.sh' -not -path "./vendor/*" | xargs shellcheck

linter_test: ## Use gometalinter check code, for local test
	@./golinter.sh "test" "${p}"
	@find . -name '*.sh' -not -path "./vendor/*" | xargs shellcheck

ineffassign:
	@ineffassign -n ${PKG_LIST_INEFFASSIGN}

race: ## Run data race detector
	@go test -race -short $(PKG_LIST)

test: ## Run unittests
	@go test -race $(PKG_LIST)

testq: ## Run unittests
	@go test $(PKG_LIST)

fmt: fmt_proto fmt_shell ## go fmt
	@go fmt ./...
	@find . -name '*.go' -not -path "./vendor/*" | xargs goimports -l -w

.PHONY: fmt_proto fmt_shell
fmt_proto: ## go fmt protobuf file
	#@find . -name '*.proto' -not -path "./vendor/*" | xargs clang-format -i

fmt_shell: ## check shell file
	@find . -name '*.sh' -not -path "./vendor/*" | xargs shfmt -w -s -i 4 -ci -bn

fmt_go: fmt_shell ## go fmt
	@go fmt ./...
	@find . -name '*.go' -not -path "./vendor/*" | xargs goimports -l -w


coverage: ## Generate global code coverage report
	@./build/tools/coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	@./build/tools/coverage.sh html;

docker: ## build docker image for bnchain run
	@sudo docker build . -f ./build/Dockerfile-run -t bnchain:latest

docker-compose: ## build docker-compose for bnchain run
	@cd build && if ! [ -d ci ]; then \
	 make -C ../ ; \
	 fi; \
	 cp bnchain* Dockerfile  docker-compose.yml *.sh ci/ && cd ci/ && ./docker-compose-pre.sh run $(proj) $(dapp)  && cd ../..

docker-compose-down: ## build docker-compose for bnchain run
	@cd build && if [ -d ci ]; then \
	 cp bnchain* Dockerfile  docker-compose* ci/ && cd ci/ && ./docker-compose-pre.sh down $(proj) $(dapp) && cd .. ; \
	 fi; \
	 cd ..

fork-test: ## build fork-test for bnchain run
	@cd build && cp bnchain* Dockerfile system-fork-test.sh docker-compose* ci/ && cd ci/ && ./docker-compose-pre.sh forktest $(proj) $(dapp) && cd ../..


clean: ## Remove previous build
	@rm -rf $(shell find . -name 'datadir' -not -path "./vendor/*")
	@rm -rf build/bnchain*
	@rm -rf build/relayd*
	@rm -rf build/*.log
	@rm -rf build/logs
	@rm -rf build/autotest/autotest
	@rm -rf build/ci
	@rm -rf build/system-rpc-test.sh
	@rm -rf tool
	@go clean

proto:protobuf

protobuf: ## Generate protbuf file of types package
#	@cd ${bnchain_PATH}/types/proto && ./create_protobuf.sh && cd ../..
	@find ./plugin/dapp -maxdepth 2 -type d  -name proto -exec make -C {} \;

depends: ## Generate depends file of types package
	@find ./plugin/dapp -maxdepth 2 -type d  -name cmd -exec make -C {} OUT="$(MKDIR)build/ci" FLAG= \;
	@find ./vendor/github.com/bnchain/bnchain/system/dapp -maxdepth 2 -type d  -name cmd -exec make -C {} OUT="$(MKDIR)build/ci" FLAG= \;

help: ## Display this help screen
	@printf "Help doc:\nUsage: make [command]\n"
	@printf "[command]\n"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

cleandata:
	rm -rf build/datadir/addrbook
	rm -rf build/datadir/blockchain.db
	rm -rf build/datadir/mavltree
	rm -rf build/bnchain.log

.PHONY: checkgofmt
checkgofmt: ## get all go files and run go fmt on them
	@files=$$(find . -name '*.go' -not -path "./vendor/*" | xargs gofmt -l -s); if [ -n "$$files" ]; then \
		  echo "Error: 'make fmt' needs to be run on:"; \
		  echo "${files}"; \
		  exit 1; \
		  fi;
	@files=$$(find . -name '*.go' -not -path "./vendor/*" | xargs goimports -l -w); if [ -n "$$files" ]; then \
		  echo "Error: 'make fmt' needs to be run on:"; \
		  echo "${files}"; \
		  exit 1; \
		  fi;

.PHONY: auto_ci_before auto_ci_after auto_ci
auto_ci_before: clean fmt protobuf
	@echo "auto_ci"
	@go version
	@protoc --version
	@mockery -version
	@docker version
	@docker-compose version
	@git version
	@git status

.PHONY: auto_ci_after
auto_ci_after: clean fmt protobuf
	@git add *.go *.sh *.proto
	@git status
	@files=$$(git status -suno);if [ -n "$$files" ]; then \
		  git status; \
		  git commit -a -m "auto ci [ci-skip]"; \
		  git push origin HEAD:$(branch); \
		  fi;

.PHONY: auto_ci
auto_fmt := find . -name '*.go' -not -path './vendor/*' | xargs goimports -l -w
auto_ci: clean fmt_proto fmt_shell protobuf
	@-find . -name '*.go' -not -path './vendor/*' | xargs gofmt -l -w -s
	@-${auto_fmt}
	@-find . -name '*.go' -not -path './vendor/*' | xargs gofmt -l -w -s
	@${auto_fmt}
	@git add -u
	@git status
	@files=$$(git status -suno); if [ -n "$$files" ]; then \
		  git add -u; \
		  git status; \
		  git commit -a -m "auto ci"; \
		  git remote add originx $(originx); \
		  git remote -v; \
		  git push --quiet --set-upstream originx HEAD:$(branch); \
		  git log -n 2; \
		  exit 1; \
		  fi;


addupstream:
	git remote add upstream https://github.com/bnchain/plugin.git
	git remote -v

sync:
	git fetch upstream
	git checkout master
	git merge upstream/master
	git push origin master

branch:
	make sync
	git checkout -b ${b}

push:
	@if [ -n "$$m" ]; then \
	git commit -a -m "${m}" ; \
	fi;
	make sync
	git checkout ${b}
	git merge master
	git push origin ${b}

pull:
	@remotelist=$$(git remote | grep ${name});if [ -z $$remotelist ]; then \
		echo ${remotelist}; \
		git remote add ${name} https://github.com/${name}/plugin.git ; \
	fi;
	git fetch ${name}
	git checkout ${name}/${b}
	git checkout -b ${name}-${b}
pullsync:
	git fetch ${name}
	git checkout ${name}-${b}
	git merge ${name}/${b}
pullpush:
	@if [ -n "$$m" ]; then \
	git commit -a -m "${m}" ; \
	fi;
	make pullsync
	git push ${name} ${name}-${b}:${b}

webhook_auto_ci: clean fmt_proto fmt_shell protobuf
	@-find . -name '*.go' -not -path './vendor/*' | xargs gofmt -l -w -s
	@-${auto_fmt}
	@-find . -name '*.go' -not -path './vendor/*' | xargs gofmt -l -w -s
	@${auto_fmt}
	@git status
	@files=$$(git status -suno);if [ -n "$$files" ]; then \
		  git status; \
		  git commit -a -m "auto ci"; \
		  git push origin ${b}; \
		  exit 0; \
		  fi;

webhook:
	git checkout ${b}
	make webhook_auto_ci name=${name} b=${b}

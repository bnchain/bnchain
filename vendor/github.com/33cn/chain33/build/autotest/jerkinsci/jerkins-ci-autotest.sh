#!/usr/bin/env bash

set -e
set -o pipefail
#set -o verbose
#set -o xtrace

# os: ubuntu16.04 x64
# first, you must install jq tool of json
# sudo apt-get install jq
# sudo apt-get install shellcheck, in order to static check shell script
# sudo apt-get install parallel
# ./run-autotest.sh build

PWD=$(cd "$(dirname "$0")" && pwd)
export PATH="$PWD:$PATH"

PROJECT_NAME="${1}"
if [[ ${PROJECT_NAME} == "" ]]; then
    PROJECT_NAME="build"
    rm -rf ../jerkinsci/tempbuild
    mv -f ./temp tempbuild
fi

NODE3="${PROJECT_NAME}_autotest_1"
CLI="docker exec ${NODE3} /root/bnchain-cli"
TEMP_CI_DIR=temp"${PROJECT_NAME}"

sedfix=""
if [ "$(uname)" == "Darwin" ]; then
    sedfix=".bak"
fi

bnchainConfig="bnchain.test.toml"
bnchainBlockTime=1
autoTestConfig="autotest.toml"
autoTestCheckTimeout=10

function config_bnchain() {

    # shellcheck disable=SC2015
    echo "# config bnchain solo test"
    # update test environment
    sed -i $sedfix 's/^Title.*/Title="local"/g' ${bnchainConfig}
    # grep -q '^TestNet' ${bnchainConfig} && sed -i $sedfix 's/^TestNet.*/TestNet=true/' ${bnchainConfig} || sed -i '/^Title/a TestNet=true' ${bnchainConfig}

    if grep -q '^TestNet' ${bnchainConfig}; then
        sed -i $sedfix 's/^TestNet.*/TestNet=true/' ${bnchainConfig}
    else
        sed -i $sedfix '/^Title/a TestNet=true' ${bnchainConfig}
    fi
}

function config_autotest() {

    echo "# config autotest"
    sed -i $sedfix 's/^checkTimeout.*/checkTimeout='${autoTestCheckTimeout}'/' ${autoTestConfig}
}

function start_bnchain() {

    # create and run docker-compose container
    docker-compose -p "${PROJECT_NAME}" -f compose-autotest.yml up --build -d

    local SLEEP=2
    sleep ${SLEEP}

    # docker-compose ps
    docker-compose -p "${PROJECT_NAME}" -f compose-autotest.yml ps

    # query node run status
    ${CLI} block last_header

    echo "=========== # save seed to wallet ============="
    result=$(${CLI} seed save -p 1314fuzamei -s "tortoise main civil member grace happy century convince father cage beach hip maid merry rib" | jq ".isok")
    if [ "${result}" = "false" ]; then
        echo "save seed to wallet error seed, result: ${result}"
        exit 1
    fi

    echo "=========== # unlock wallet ============="
    result=$(${CLI} wallet unlock -p 1314fuzamei -t 0 | jq ".isok")
    if [ "${result}" = "false" ]; then
        exit 1
    fi

    echo "=========== # import private key returnAddr ============="
    result=$(${CLI} account import_key -k CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944 -l returnAddr | jq ".label")
    echo "${result}"
    if [ -z "${result}" ]; then
        exit 1
    fi

    echo "=========== # import private key mining ============="
    result=$(${CLI} account import_key -k 4257D8692EF7FE13C68B65D6A52F03933DB2FA5CE8FAF210B5B8B80C721CED01 -l minerAddr | jq ".label")
    echo "${result}"
    if [ -z "${result}" ]; then
        exit 1
    fi

    echo "=========== #transfer to miner addr ============="
    hash=$(${CLI} send coins transfer -a 10000 -n test -t 12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv -k CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944)
    echo "${hash}"
    sleep ${bnchainBlockTime}
    result=$(${CLI} tx query -s "${hash}" | jq '.receipt.tyName')
    if [[ ${result} != '"ExecOk"' ]]; then
        echo "Failed"
        ${CLI} tx query -s "${hash}" | jq '.' | cat
        exit 1
    fi

    echo "=========== #transfer to token amdin ============="
    hash=$(${CLI} send coins transfer -a 10 -n test -t 1Q8hGLfoGe63efeWa8fJ4Pnukhkngt6poK -k CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944)

    echo "${hash}"
    sleep ${bnchainBlockTime}
    result=$(${CLI} tx query -s "${hash}" | jq '.receipt.tyName')
    if [[ ${result} != '"ExecOk"' ]]; then
        echo "Failed"
        ${CLI} tx query -s "${hash}" | jq '.' | cat
        exit 1
    fi

    echo "=========== #config token blacklist ============="
    rawData=$(${CLI} config config_tx -c token-blacklist -o add -v BTC)
    signData=$(${CLI} wallet sign -d "${rawData}" -k 0xc34b5d9d44ac7b754806f761d3d4d2c4fe5214f6b074c19f069c4f5c2a29c8cc)
    hash=$(${CLI} wallet send -d "${signData}")

    echo "${hash}"
    sleep ${bnchainBlockTime}
    result=$(${CLI} tx query -s "${hash}" | jq '.receipt.tyName')
    if [[ ${result} != '"ExecOk"' ]]; then
        echo "Failed"
        ${CLI} tx query -s "${hash}" | jq '.' | cat
        exit 1
    fi

}

function start_autotest() {

    echo "=========== #start auto-test program ============="
    docker exec "${NODE3}" /root/autotest

}

function stop_bnchain() {

    rv=$?
    echo "=========== #stop docker-compose ============="
    docker-compose -p "${PROJECT_NAME}" -f compose-autotest.yml down
    echo "=========== #remove related images ============"
    docker rmi "${PROJECT_NAME}"_autotest || true
    cd ../ && rm -rf ./"${TEMP_CI_DIR}"
    exit ${rv}
}

function main() {

    cd "${TEMP_CI_DIR}" && cp ../compose-autotest.yml ../Dockerfile-autotest ./
    config_bnchain
    config_autotest
    start_bnchain
    start_autotest

}

#trap exit
trap "stop_bnchain" INT TERM EXIT

# run script
main

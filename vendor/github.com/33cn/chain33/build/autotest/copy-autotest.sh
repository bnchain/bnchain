#!/usr/bin/env bash

set -e
set -o pipefail
#set -o verbose
#set -o xtrace

# os: ubuntu16.04 x64

#bnchain dapp autotest root directory
declare -a bnchainAutoTestDirs=("system" "plugin" "vendor/github.com/bnchain/bnchain/system" "vendor/github.com/bnchain/plugin/plugin")

#copy auto test to specific directory
# check args
if [ "$#" -lt 1 ]; then
    echo "Usage: $0 directory list"
    exit 1
fi

function copyAutoTestConfig() {

    echo "#copy auto test config to path \"$1\""
    local AutoTestConfigFile="$1/autotest.toml"

    #pre config auto test
    {

        echo 'cliCmd="./bnchain-cli"'
        echo "checkTimeout=60"
    } >"${AutoTestConfigFile}"

    #copy all the dapp test case config file
    for rootDir in "${bnchainAutoTestDirs[@]}"; do

        if [ ! -d ../../"${rootDir}" ]; then
            continue
        fi

        testDirArr=$(find ../../"${rootDir}" -type d -name autotest)

        for autotest in ${testDirArr}; do

            dapp=$(basename "$(dirname "${autotest}")")
            dappConfig=${autotest}/${dapp}.toml

            #make sure dapp have auto test config
            if [ -e "${dappConfig}" ]; then

                cp "${dappConfig}" "$1"/

                #add dapp test case config
                {
                    echo "[[TestCaseFile]]"
                    echo "dapp=\"$dapp\""
                    echo "filename=\"$dapp.toml\""
                } >>"${AutoTestConfigFile}"

            fi

        done
    done
}

function copybnchain() {

    echo "# copy bnchain bin to path \"$1\", make sure build bnchain"
    cp ../bnchain ../bnchain-cli ../bnchain.toml "$1"
    find ../../ -path '*cmd/bnchain/bnchain.test.toml' -exec cp {} "$1" ';'
}

for dir in "$@"; do

    #check dir exist
    if [ ! -d "${dir}" ]; then
        mkdir "${dir}"
    fi
    cp autotest "${dir}"
    copyAutoTestConfig "${dir}"
    copybnchain "${dir}"
    echo "# all copy have done!"

done

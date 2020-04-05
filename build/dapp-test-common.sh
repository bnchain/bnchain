#!/usr/bin/env bash

RAW_TX_HASH=""
LAST_BLOCK_HASH=""
CASE_ERR=""

#color
RED='\033[1;31m'
GRE='\033[1;32m'
NOC='\033[0m'

echo_rst() {
    if [ "$2" -eq 0 ]; then
        echo -e "${GRE}$1 ok${NOC}"
    elif [ "$2" -eq 2 ]; then
        echo -e "${GRE}$1 not support${NOC}"
        CASE_ERR="err"
        echo $CASE_ERR
    else
        echo -e "${RED}$1 fail${NOC}"
        CASE_ERR="err"
        echo $CASE_ERR
    fi
}

bnchain_BlockWait() {
    local MAIN_HTTP=$2
    local req='"method":"bnchain.GetLastHeader","params":[]'

    cur_height=$(curl -ksd "{$req}" "${MAIN_HTTP}" | jq ".result.height")
    expect=$((cur_height + ${1}))

    local count=0
    while true; do
        new_height=$(curl -ksd "{$req}" "${MAIN_HTTP}" | jq ".result.height")
        if [ "${new_height}" -ge "${expect}" ]; then
            break
        fi
        count=$((count + 1))
        sleep 1
    done
    echo "wait new block $count s, cur height=$expect,old=$cur_height"
}

bnchain_QueryTx() {
    local MAIN_HTTP=$2
    bnchain_BlockWait 1 "$MAIN_HTTP"
    local txhash="$1"
    local req='"method":"bnchain.QueryTransaction","params":[{"hash":"'"$txhash"'"}]'

    local times=10
    while true; do
        ret=$(curl -ksd "{$req}" "${MAIN_HTTP}" | jq -r ".result.tx.hash")
        if [ "${ret}" != "${1}" ]; then
            bnchain_BlockWait 1 "$MAIN_HTTP"
            times=$((times - 1))
            if [ $times -le 0 ]; then
                echo "====query tx=$1 failed"
                curl -ksd "{$req}" "${MAIN_HTTP}"
                exit 1
            fi
        else
            RAW_TX_HASH=$txhash
            echo "====query tx=$RAW_TX_HASH success"
            break
        fi
    done
}

bnchain_SendTx() {
    local signedTx=$1
    local MAIN_HTTP=$2

    req='"method":"bnchain.SendTransaction","params":[{"token":"BTY","data":"'"$signedTx"'"}]'
    resp=$(curl -ksd "{$req}" "${MAIN_HTTP}")
    err=$(jq '(.error)' <<<"$resp")
    txhash=$(jq -r ".result" <<<"$resp")

    if [ "$err" == null ]; then
        bnchain_QueryTx "$txhash" "$MAIN_HTTP"
    else
        echo "send tx error:$err"
    fi
}

bnchain_SendToAddress() {
    local from="$1"
    local to="$2"
    local amount=$3
    local MAIN_HTTP=$4

    local req='"method":"bnchain.SendToAddress", "params":[{"from":"'"$from"'","to":"'"$to"'", "amount":'"$amount"', "note":"test\n"}]'
    resp=$(curl -ksd "{$req}" "${MAIN_HTTP}")
    ok=$(jq '(.error|not) and (.result.hash|length==66)' <<<"$resp")

    [ "$ok" == true ]

    hash=$(jq -r ".result.hash" <<<"$resp")
    bnchain_QueryTx "$hash" "$MAIN_HTTP"
}

bnchain_ImportPrivkey() {
    local pri="$1"
    local acc="$2"
    local label="$3"
    local MAIN_HTTP=$4

    local req='"method":"bnchain.ImportPrivkey", "params":[{"privkey":"'"$pri"'", "label":"'"$label"'"}]'
    resp=$(curl -ksd "{$req}" "$MAIN_HTTP")
    ok=$(jq '(((.error|not) and (.result.label=="'"$label"'") and (.result.acc.addr == "'"$acc"'")) or (.error=="ErrPrivkeyExist"))' <<<"$resp")

    [ "$ok" == true ]
}

bnchain_SignRawTx() {
    local txHex="$1"
    local priKey="$2"
    local MAIN_HTTP=$3

    local req='"method":"bnchain.SignRawTx","params":[{"privkey":"'"$priKey"'","txHex":"'"$txHex"'","expire":"120s"}]'
    signedTx=$(curl -ksd "{$req}" "${MAIN_HTTP}" | jq -r ".result")

    if [ "$signedTx" != null ]; then
        bnchain_SendTx "$signedTx" "${MAIN_HTTP}"
    else
        echo "signedTx null error"
    fi
}

bnchain_QueryBalance() {
    local addr=$1
    local MAIN_HTTP=$2
    req='"method":"bnchain.GetAllExecBalance","params":[{"addr":"'"${addr}"'"}]'
    #echo "#request: $req"
    resp=$(curl -ksd "{$req}" "${MAIN_HTTP}")
    echo "#response: $resp"
    ok=$(jq '(.error|not) and (.result != "")' <<<"$resp")
    [ "$ok" == true ]

    echo "$resp" | jq -r ".result"
}

bnchain_QueryExecBalance() {
    local addr=$1
    local exec=$2
    local MAIN_HTTP=$3

    req='{"method":"bnchain.GetBalance", "params":[{"addresses" : ["'"${addr}"'"], "execer" : "'"${exec}"'"}]}'
    resp=$(curl -ksd "$req" "${MAIN_HTTP}")
    echo "#response: $resp"
    ok=$(jq '(.error|not) and (.result[0] | [has("balance", "frozen"), true] | unique | length == 1)' <<<"$resp")
    [ "$ok" == true ]
}

bnchain_GetAccounts() {
    local MAIN_HTTP=$1
    resp=$(curl -ksd '{"jsonrpc":"2.0","id":2,"method":"bnchain.GetAccounts","params":[{}]}' -H 'content-type:text/plain;' "${MAIN_HTTP}")
    echo "$resp"
}

bnchain_LastBlockhash() {
    local MAIN_HTTP=$1
    result=$(curl -ksd '{"method":"bnchain.GetLastHeader","params":[{}]}' -H 'content-type:text/plain;' "${MAIN_HTTP}" | jq -r ".result.hash")
    LAST_BLOCK_HASH=$result
    echo -e "######\\n  last blockhash is $LAST_BLOCK_HASH  \\n######"
}

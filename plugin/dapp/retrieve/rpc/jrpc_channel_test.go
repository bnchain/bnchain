// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc_test

import (
	"strings"
	"testing"

	commonlog "github.com/bnchain/bnchain/common/log"
	"github.com/bnchain/bnchain/rpc/jsonclient"
	rpctypes "github.com/bnchain/bnchain/rpc/types"
	"github.com/bnchain/bnchain/types"
	"github.com/bnchain/bnchain/util/testnode"
	"github.com/bnchain/plugin/plugin/dapp/retrieve/rpc"
	pty "github.com/bnchain/plugin/plugin/dapp/retrieve/types"
	"github.com/stretchr/testify/assert"

	_ "github.com/bnchain/bnchain/system"
	_ "github.com/bnchain/plugin/plugin"
)

func init() {
	commonlog.SetLogLevel("error")
}

func TestJRPCChannel(t *testing.T) {
	// 启动RPCmocker
	mocker := testnode.New("--notset--", nil)
	defer func() {
		mocker.Close()
	}()
	mocker.Listen()

	jrpcClient := mocker.GetJSONC()

	testCases := []struct {
		fn func(*testing.T, *jsonclient.JSONClient) error
	}{
		{fn: testBackupCmd},
		{fn: testPrepareCmd},
		{fn: testPerformCmd},
		{fn: testCancelCmd},
		{fn: testRetrieveQueryCmd},
	}
	for index, testCase := range testCases {
		err := testCase.fn(t, jrpcClient)
		if err == nil {
			continue
		}
		assert.NotEqualf(t, err, types.ErrActionNotSupport, "test index %d", index)
		if strings.Contains(err.Error(), "rpc: can't find") {
			assert.FailNowf(t, err.Error(), "test index %d", index)
		}
	}
}

func testBackupCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := rpc.RetrieveBackupTx{}
	return jrpc.Call("retrieve.CreateRawRetrieveBackupTx", params, nil)
}

func testPrepareCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := rpc.RetrievePrepareTx{}
	return jrpc.Call("retrieve.CreateRawRetrievePrepareTx", params, nil)
}

func testPerformCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := rpc.RetrievePerformTx{}
	return jrpc.Call("retrieve.CreateRawRetrievePerformTx", params, nil)
}

func testCancelCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := rpc.RetrieveCancelTx{}
	return jrpc.Call("retrieve.CreateRawRetrieveCancelTx", params, nil)
}

func testRetrieveQueryCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var rep interface{}
	var params rpctypes.Query4Jrpc
	req := &pty.ReqRetrieveInfo{}
	params.Execer = "retrieve"
	params.FuncName = "GetRetrieveInfo"
	params.Payload = types.MustPBToJSON(req)
	rep = &pty.RetrieveQuery{}
	return jrpc.Call("bnchain.Query", params, rep)
}

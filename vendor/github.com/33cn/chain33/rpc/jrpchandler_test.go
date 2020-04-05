// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"errors"
	"testing"

	"encoding/hex"

	"github.com/bnchain/bnchain/account"
	"github.com/bnchain/bnchain/client/mocks"
	"github.com/bnchain/bnchain/common"
	rpctypes "github.com/bnchain/bnchain/rpc/types"
	_ "github.com/bnchain/bnchain/system"
	cty "github.com/bnchain/bnchain/system/dapp/coins/types"
	mty "github.com/bnchain/bnchain/system/dapp/manage/types"
	"github.com/bnchain/bnchain/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDecodeLogErr(t *testing.T) {
	enc := "0001020304050607"
	dec := []byte{0, 1, 2, 3, 4, 5, 6, 7}

	hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogErr,
		Log: "0x" + enc,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   1,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogErr", result.Logs[0].TyName)
	assert.Equal(t, int32(types.TyLogErr), result.Logs[0].Ty)
}

func TestDecodeLogFee(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogFee,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogFee", result.Logs[0].TyName)
}

func TestDecodeLogTransfer(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogTransfer,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogTransfer", result.Logs[0].TyName)
}

func TestDecodeLogGenesis(t *testing.T) {
	enc := "0001020304050607"

	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogGenesis,
		Log: "0x" + enc,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	//这个已经废弃
	assert.Equal(t, "unkownType", result.Logs[0].TyName)
}

func TestDecodeLogDeposit(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogDeposit,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogDeposit", result.Logs[0].TyName)
}

func TestDecodeLogExecTransfer(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptExecAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogExecTransfer,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogExecTransfer", result.Logs[0].TyName)
}

func TestDecodeLogExecWithdraw(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptExecAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogExecWithdraw,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogExecWithdraw", result.Logs[0].TyName)
}

func TestDecodeLogExecDeposit(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptExecAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogExecDeposit,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogExecDeposit", result.Logs[0].TyName)
}

func TestDecodeLogExecFrozen(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptExecAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogExecFrozen,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogExecFrozen", result.Logs[0].TyName)
}

func TestDecodeLogExecActive(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptExecAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogExecActive,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogExecActive", result.Logs[0].TyName)
}

func TestDecodeLogGenesisTransfer(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogGenesisTransfer,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogGenesisTransfer", result.Logs[0].TyName)
}

func TestDecodeLogGenesisDeposit(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogGenesisDeposit,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogGenesisDeposit", result.Logs[0].TyName)
}

func TestDecodeLogModifyConfig(t *testing.T) {
	var logTmp = &types.ReceiptConfig{}
	dec := types.Encode(logTmp)
	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  mty.TyLogModifyConfig,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("manage"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogModifyConfig", result.Logs[0].TyName)
}

func newTestbnchain(api *mocks.QueueProtocolAPI) *bnchain {
	return &bnchain{
		cli: channelClient{
			QueueProtocolAPI: api,
			accountdb:        account.NewCoinsAccount(),
		},
	}
}

func Testbnchain_CreateRawTransaction(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	// var result interface{}
	// api.On("CreateRawTransaction", nil, &result).Return()
	testbnchain := newTestbnchain(api)
	var testResult interface{}
	err := testbnchain.CreateRawTransaction(nil, &testResult)
	assert.Nil(t, testResult)
	assert.NotNil(t, err)

	tx := &rpctypes.CreateTx{
		To:          "184wj4nsgVxKyz2NhM3Yb5RK5Ap6AFRFq2",
		Amount:      10,
		Fee:         1,
		Note:        "12312",
		IsWithdraw:  false,
		IsToken:     false,
		TokenSymbol: "",
		ExecName:    types.ExecName("coins"),
	}

	err = testbnchain.CreateRawTransaction(tx, &testResult)
	assert.NotNil(t, testResult)
	assert.Nil(t, err)
}

func Testbnchain_ReWriteRawTx(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)
	txHex1 := "0a05636f696e73122c18010a281080c2d72f222131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b7120a08d0630a696c0b3f78dd9ec083a2131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b71"
	//txHex2 := "0a05636f696e73122d18010a29108084af5f222231484c53426e7437486e486a7857797a636a6f573863663259745550663337594d6320a08d0630dbc4cbf6fbc4e1d0533a2231484c53426e7437486e486a7857797a636a6f573863663259745550663337594d63"

	reTx := &rpctypes.ReWriteRawTx{
		Tx:     txHex1,
		Fee:    29977777777,
		Expire: "130s",
		To:     "aabbccdd",
		Index:  0,
	}
	var testResult interface{}
	err := testbnchain.ReWriteRawTx(reTx, &testResult)
	assert.Nil(t, err)
	assert.NotNil(t, testResult)
	assert.NotEqual(t, txHex1, testResult)
	txData, err := hex.DecodeString(testResult.(string))
	assert.Nil(t, err)
	tx := &types.Transaction{}
	err = types.Decode(txData, tx)
	assert.Nil(t, err)
	assert.Equal(t, tx.Fee, reTx.Fee)
	assert.Equal(t, reTx.To, tx.To)

}

func Testbnchain_CreateTxGroup(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)
	var testResult interface{}
	err := testbnchain.CreateRawTxGroup(nil, &testResult)
	assert.Nil(t, testResult)
	assert.NotNil(t, err)

	txHex1 := "0a05636f696e73122c18010a281080c2d72f222131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b7120a08d0630a696c0b3f78dd9ec083a2131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b71"
	txHex2 := "0a05636f696e73122d18010a29108084af5f222231484c53426e7437486e486a7857797a636a6f573863663259745550663337594d6320a08d0630dbc4cbf6fbc4e1d0533a2231484c53426e7437486e486a7857797a636a6f573863663259745550663337594d63"
	txs := &types.CreateTransactionGroup{
		Txs: []string{txHex1, txHex2},
	}
	err = testbnchain.CreateRawTxGroup(txs, &testResult)
	assert.Nil(t, err)
	tx, err := decodeTx(testResult.(string))
	assert.Nil(t, err)
	tg, err := tx.GetTxGroup()
	assert.Nil(t, err)
	if len(tg.GetTxs()) != 2 {
		t.Error("Test createtxgroup failed")
		return
	}
	err = tx.Check(0, types.GInt("MinFee"), types.GInt("MaxFee"))
	assert.Nil(t, err)
}

func Testbnchain_SendTransaction(t *testing.T) {
	if types.IsPara() {
		t.Skip()
		return
	}
	api := new(mocks.QueueProtocolAPI)
	tx := &types.Transaction{}
	api.On("SendTx", tx).Return(nil, errors.New("error value"))
	testbnchain := newTestbnchain(api)
	var testResult interface{}
	data := rpctypes.RawParm{
		Data: "",
	}
	err := testbnchain.SendTransaction(data, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetHexTxByHash(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	api.On("QueryTx", &types.ReqHash{Hash: []byte("")}).Return(nil, errors.New("error value"))
	testbnchain := newTestbnchain(api)
	var testResult interface{}
	data := rpctypes.QueryParm{
		Hash: "",
	}
	err := testbnchain.GetHexTxByHash(data, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_QueryTransaction(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	api.On("QueryTx", &types.ReqHash{Hash: []byte("")}).Return(nil, errors.New("error value"))
	testbnchain := newTestbnchain(api)
	var testResult interface{}
	data := rpctypes.QueryParm{
		Hash: "",
	}
	err := testbnchain.QueryTransaction(data, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_QueryTransactionOk(t *testing.T) {
	data := rpctypes.QueryParm{
		Hash: "",
	}
	var act = &cty.CoinsAction{
		Ty: 1,
	}
	payload := types.Encode(act)
	var tx = &types.Transaction{
		Execer:  []byte(types.ExecName("ticket")),
		Payload: payload,
	}

	var logTmp = &types.ReceiptAccountTransfer{}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	strdec = "0x" + strdec

	rlog := &types.ReceiptLog{
		Ty:  types.TyLogTransfer,
		Log: []byte(strdec),
	}

	logs := []*types.ReceiptLog{}
	logs = append(logs, rlog)

	var rdata = &types.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	reply := types.TransactionDetail{
		Tx:      tx,
		Receipt: rdata,
		Height:  10,
	}

	api := new(mocks.QueueProtocolAPI)
	api.On("QueryTx", &types.ReqHash{Hash: []byte("")}).Return(&reply, nil)
	testbnchain := newTestbnchain(api)
	var testResult interface{}

	err := testbnchain.QueryTransaction(data, &testResult)
	t.Log(err)
	assert.Nil(t, err)
	assert.Equal(t, testResult.(*rpctypes.TransactionDetail).Height, reply.Height)
	assert.Equal(t, testResult.(*rpctypes.TransactionDetail).Tx.Execer, string(tx.Execer))

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetBlocks(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	api.On("GetBlocks", &types.ReqBlocks{Pid: []string{""}}).Return(&types.BlockDetails{Items: []*types.BlockDetail{{}}}, nil)
	testbnchain := newTestbnchain(api)
	var testResult interface{}
	data := rpctypes.BlockParam{}
	err := testbnchain.GetBlocks(data, &testResult)
	t.Log(err)
	assert.NoError(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetLastHeader(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	api.On("GetLastHeader", mock.Anything).Return(&types.Header{}, nil)
	testbnchain := newTestbnchain(api)
	var testResult interface{}
	data := &types.ReqNil{}
	err := testbnchain.GetLastHeader(data, &testResult)
	t.Log(err)
	assert.NotNil(t, &testResult)
	assert.NoError(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetTxByAddr(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	api.On("GetTransactionByAddr", mock.Anything).Return(&types.ReplyTxInfos{TxInfos: []*types.ReplyTxInfo{{}}}, nil)
	var testResult interface{}
	data := types.ReqAddr{}
	err := testbnchain.GetTxByAddr(data, &testResult)
	t.Log(err)
	assert.NotNil(t, testResult)
	assert.NoError(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetTxByHashes(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	api.On("GetTransactionByHash", mock.Anything).Return(&types.TransactionDetails{}, nil)
	var testResult interface{}
	data := rpctypes.ReqHashes{}
	data.Hashes = append(data.Hashes, "0xdcf13a93e3bf58534c773e13d339894c18dafbd3ff273a9d1caa0c2bec8e8cd6")
	err := testbnchain.GetTxByHashes(data, &testResult)
	t.Log(err)
	assert.NotNil(t, testResult)
	assert.NoError(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetMempool(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	api.On("GetMempool").Return(&types.ReplyTxList{Txs: []*types.Transaction{{}}}, nil)
	var testResult interface{}
	data := &types.ReqNil{}
	err := testbnchain.GetMempool(data, &testResult)
	t.Log(err)
	assert.NotNil(t, testResult)
	assert.NoError(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetAccountsV2(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	api.On("WalletGetAccountList", mock.Anything).Return(&types.WalletAccounts{Wallets: []*types.WalletAccount{{}}}, nil)
	var testResult interface{}
	err := testbnchain.GetAccountsV2(nil, &testResult)
	t.Log(err)
	assert.NotNil(t, testResult)
	assert.NoError(t, err)
}

func Testbnchain_GetAccounts(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	api.On("WalletGetAccountList", mock.Anything).Return(nil, errors.New("error value"))
	var testResult interface{}
	data := &types.ReqAccountList{}
	err := testbnchain.GetAccounts(data, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_NewAccount(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	api.On("NewAccount", &types.ReqNewAccount{}).Return(nil, errors.New("error value"))

	var testResult interface{}
	err := testbnchain.NewAccount(types.ReqNewAccount{}, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_WalletTxList(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.ReqWalletTransactionList{FromTx: []byte("")}
	api.On("WalletTransactionList", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := rpctypes.ReqWalletTransactionList{}
	err := testbnchain.WalletTxList(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_ImportPrivkey(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.ReqWalletImportPrivkey{}
	api.On("WalletImportprivkey", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqWalletImportPrivkey{}
	err := testbnchain.ImportPrivkey(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_SendToAddress(t *testing.T) {
	if types.IsPara() {
		t.Skip()
		return
	}
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.ReqWalletSendToAddress{}
	api.On("WalletSendToAddress", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqWalletSendToAddress{}
	err := testbnchain.SendToAddress(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_SetTxFee(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.ReqWalletSetFee{}
	api.On("WalletSetFee", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqWalletSetFee{}
	err := testbnchain.SetTxFee(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_SetLabl(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.ReqWalletSetLabel{}
	api.On("WalletSetLabel", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqWalletSetLabel{}
	err := testbnchain.SetLabl(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_MergeBalance(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.ReqWalletMergeBalance{}
	api.On("WalletMergeBalance", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqWalletMergeBalance{}
	err := testbnchain.MergeBalance(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_SetPasswd(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.ReqWalletSetPasswd{}
	api.On("WalletSetPasswd", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqWalletSetPasswd{}
	err := testbnchain.SetPasswd(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_Lock(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	// expected := types.ReqNil{}
	api.On("WalletLock").Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqNil{}
	err := testbnchain.Lock(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_UnLock(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.WalletUnLock{}
	api.On("WalletUnLock", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.WalletUnLock{}
	err := testbnchain.UnLock(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetPeerInfo(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	api.On("PeerInfo").Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqNil{}
	err := testbnchain.GetPeerInfo(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetPeerInfoOk(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	var peerlist types.PeerList
	var pr = &types.Peer{
		Addr: "abcdsd",
	}
	peerlist.Peers = append(peerlist.Peers, pr)

	api.On("PeerInfo").Return(&peerlist, nil)
	var testResult interface{}
	var in types.ReqNil
	_ = testbnchain.GetPeerInfo(in, &testResult)
	assert.Equal(t, testResult.(*rpctypes.PeerList).Peers[0].Addr, peerlist.Peers[0].Addr)
}

func Testbnchain_GetHeaders(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.ReqBlocks{}
	api.On("GetHeaders", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqBlocks{}
	err := testbnchain.GetHeaders(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetHeadersOk(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	var headers types.Headers
	var header = &types.Header{
		TxCount: 10,
	}
	headers.Items = append(headers.Items, header)

	expected := &types.ReqBlocks{}
	api.On("GetHeaders", expected).Return(&headers, nil)

	var testResult interface{}
	actual := types.ReqBlocks{}
	err := testbnchain.GetHeaders(actual, &testResult)
	assert.Nil(t, err)
	assert.Equal(t, testResult.(*rpctypes.Headers).Items[0].TxCount, header.TxCount)

}

func Testbnchain_GetLastMemPool(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	// expected := &types.ReqBlocks{}
	api.On("GetLastMempool").Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqNil{}
	err := testbnchain.GetLastMemPool(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetProperFee(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := types.ReqProperFee{}
	api.On("GetProperFee", &expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	err := testbnchain.GetProperFee(expected, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetBlockOverview(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.ReqHash{Hash: []byte{}}
	api.On("GetBlockOverview", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := rpctypes.QueryParm{}
	err := testbnchain.GetBlockOverview(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetBlockOverviewOk(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)
	var head = &types.Header{
		Hash: []byte("123456"),
	}
	var replyblock = &types.BlockOverview{
		Head:    head,
		TxCount: 1,
	}

	expected := &types.ReqHash{Hash: []byte{0x12, 0x34, 0x56}}
	api.On("GetBlockOverview", expected).Return(replyblock, nil)

	var testResult interface{}
	actual := rpctypes.QueryParm{Hash: "123456"}

	err := testbnchain.GetBlockOverview(actual, &testResult)
	t.Log(err)
	assert.Nil(t, err)
	assert.Equal(t, testResult.(*rpctypes.BlockOverview).TxCount, replyblock.TxCount)
	assert.Equal(t, testResult.(*rpctypes.BlockOverview).Head.Hash, common.ToHex(replyblock.Head.Hash))
	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetAddrOverview(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.ReqAddr{}
	api.On("GetAddrOverview", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqAddr{}
	err := testbnchain.GetAddrOverview(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	// mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetBlockHash(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.ReqInt{}
	api.On("GetBlockHash", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqInt{}
	err := testbnchain.GetBlockHash(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GenSeed(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.GenSeedLang{}
	api.On("GenSeed", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.GenSeedLang{}
	err := testbnchain.GenSeed(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_SaveSeed(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.SaveSeedByPw{}
	api.On("SaveSeed", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.SaveSeedByPw{}
	err := testbnchain.SaveSeed(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetSeed(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	expected := &types.GetSeedByPw{}
	api.On("GetSeed", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.GetSeedByPw{}
	err := testbnchain.GetSeed(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func Testbnchain_GetWalletStatus(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)

	api.On("GetWalletStatus").Return(nil, errors.New("error value")).Once()

	var testResult interface{}
	actual := types.ReqNil{}
	err := testbnchain.GetWalletStatus(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	expect := types.WalletStatus{
		IsWalletLock: true,
		IsAutoMining: true,
		IsHasSeed:    false,
		IsTicketLock: false,
	}
	api.On("GetWalletStatus").Return(&expect, nil).Once()
	err = testbnchain.GetWalletStatus(actual, &testResult)
	t.Log(err)
	assert.Nil(t, err)
	status, ok := testResult.(*rpctypes.WalletStatus)
	if !ok {
		t.Error("GetWalletStatus type error")
	}
	assert.Equal(t, expect.IsWalletLock, status.IsWalletLock)
	assert.Equal(t, expect.IsAutoMining, status.IsAutoMining)
	assert.Equal(t, expect.IsHasSeed, status.IsHasSeed)
	assert.Equal(t, expect.IsTicketLock, status.IsTicketLock)

	mock.AssertExpectationsForObjects(t, api)
}

// func Testbnchain_GetBalance(t *testing.T) {
// 	api := new(mocks.QueueProtocolAPI)
// 	testbnchain := newTestbnchain(api)
//
// 	expected := &types.ReqBalance{}
// 	api.On("GetBalance",expected).Return(nil, errors.New("error value"))
//
// 	var testResult interface{}
// 	actual := types.ReqBalance{}
// 	err := testbnchain.GetBalance(actual, &testResult)
// 	t.Log(err)
// 	assert.Equal(t, nil, testResult)
// 	assert.NotNil(t, err)
//
// 	mock.AssertExpectationsForObjects(t, api)
// }

// ----------------------------

func Testbnchain_Version(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testbnchain := newTestbnchain(api)
	var testResult interface{}
	in := &types.ReqNil{}
	ver := &types.VersionInfo{bnchain: "6.0.2"}
	api.On("Version", mock.Anything).Return(ver, nil)
	err := testbnchain.Version(in, &testResult)
	t.Log(err)
	t.Log(testResult)
	assert.Equal(t, nil, err)
	assert.NotNil(t, testResult)
}

func Testbnchain_GetTimeStatus(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var result interface{}
	err := client.GetTimeStatus(&types.ReqNil{}, &result)
	assert.Nil(t, err)
}

func Testbnchain_GetLastBlockSequence(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var result interface{}
	api.On("GetLastBlockSequence", mock.Anything).Return(nil, types.ErrInvalidParam)
	err := client.GetLastBlockSequence(&types.ReqNil{}, &result)
	assert.NotNil(t, err)

	api = new(mocks.QueueProtocolAPI)
	client = newTestbnchain(api)
	var result2 interface{}
	lastSeq := types.Int64{Data: 1}
	api.On("GetLastBlockSequence", mock.Anything).Return(&lastSeq, nil)
	err = client.GetLastBlockSequence(&types.ReqNil{}, &result2)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), result2)
}

func Testbnchain_GetBlockSequences(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var result interface{}
	api.On("GetBlockSequences", mock.Anything).Return(nil, types.ErrInvalidParam)
	err := client.GetBlockSequences(rpctypes.BlockParam{}, &result)
	assert.NotNil(t, err)

	api = new(mocks.QueueProtocolAPI)
	client = newTestbnchain(api)
	var result2 interface{}
	blocks := types.BlockSequences{}
	blocks.Items = make([]*types.BlockSequence, 0)
	blocks.Items = append(blocks.Items, &types.BlockSequence{Hash: []byte("h1"), Type: 1})
	api.On("GetBlockSequences", mock.Anything).Return(&blocks, nil)
	err = client.GetBlockSequences(rpctypes.BlockParam{}, &result2)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result2.(*rpctypes.ReplyBlkSeqs).BlkSeqInfos))
}

func Testbnchain_GetBlockByHashes(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var testResult interface{}
	in := rpctypes.ReqHashes{Hashes: []string{}}
	in.Hashes = append(in.Hashes, common.ToHex([]byte("h1")))
	api.On("GetBlockByHashes", mock.Anything).Return(&types.BlockDetails{}, nil)
	err := client.GetBlockByHashes(in, &testResult)
	assert.Nil(t, err)

	api = new(mocks.QueueProtocolAPI)
	client = newTestbnchain(api)
	var testResult2 interface{}
	api.On("GetBlockByHashes", mock.Anything).Return(nil, types.ErrInvalidParam)
	err = client.GetBlockByHashes(in, &testResult2)
	assert.NotNil(t, err)
}

func Testbnchain_CreateTransaction(t *testing.T) {
	client := newTestbnchain(nil)

	var result interface{}
	err := client.CreateTransaction(nil, &result)
	assert.NotNil(t, err)

	in := &rpctypes.CreateTxIn{Execer: "notExist", ActionName: "x", Payload: []byte("x")}
	err = client.CreateTransaction(in, &result)
	assert.Equal(t, types.ErrExecNotFound, err)

	in = &rpctypes.CreateTxIn{Execer: types.ExecName("coins"), ActionName: "notExist", Payload: []byte("x")}
	err = client.CreateTransaction(in, &result)
	assert.Equal(t, types.ErrActionNotSupport, err)

	in = &rpctypes.CreateTxIn{
		Execer:     types.ExecName("coins"),
		ActionName: "Transfer",
		Payload:    []byte("{\"to\": \"addr\", \"amount\":\"10\"}"),
	}
	err = client.CreateTransaction(in, &result)
	assert.Nil(t, err)
}

func Testbnchain_GetExecBalance(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var testResult interface{}
	in := &types.ReqGetExecBalance{}
	api.On("StoreList", mock.Anything).Return(&types.StoreListReply{}, nil)
	err := client.GetExecBalance(in, &testResult)
	assert.Nil(t, err)

	api = new(mocks.QueueProtocolAPI)
	client = newTestbnchain(api)
	var testResult2 interface{}
	api.On("StoreList", mock.Anything).Return(nil, types.ErrInvalidParam)
	err = client.GetExecBalance(in, &testResult2)
	assert.NotNil(t, err)
}

func Testbnchain_GetBalance(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)

	var addrs = []string{"1Jn2qu84Z1SUUosWjySggBS9pKWdAP3tZt"}
	cases := []struct {
		In types.ReqBalance
	}{
		{In: types.ReqBalance{
			Execer:    types.ExecName("coins"),
			Addresses: addrs,
		}},
		{In: types.ReqBalance{
			Execer:    types.ExecName("ticket"),
			Addresses: addrs,
		}},

		{In: types.ReqBalance{
			AssetSymbol: "bty",
			AssetExec:   "coins",
			Execer:      types.ExecName("ticket"),
			Addresses:   addrs,
		}},
		{In: types.ReqBalance{
			AssetSymbol: "bty",
			AssetExec:   "coins",
			Execer:      types.ExecName("coins"),
			Addresses:   addrs,
		}},
	}

	for _, c := range cases {
		c := c
		t.Run("test GetBalance", func(t *testing.T) {
			head := &types.Header{StateHash: []byte("sdfadasds")}
			api.On("GetLastHeader").Return(head, nil)

			var acc = &types.Account{Addr: "1Jn2qu84Z1SUUosWjySggBS9pKWdAP3tZt", Balance: 100}
			accv := types.Encode(acc)
			storevalue := &types.StoreReplyValue{}
			storevalue.Values = append(storevalue.Values, accv)
			api.On("StoreGet", mock.Anything).Return(storevalue, nil)

			var data interface{}
			err := client.GetBalance(c.In, &data)
			assert.Nil(t, err)
			result := data.([]*rpctypes.Account)
			assert.Equal(t, 1, len(result))
			//t.Error("result", "x", result)
			assert.Equal(t, acc.Addr, result[0].Addr)
			assert.Equal(t, int64(100), result[0].Balance)
		})
	}
}

func Testbnchain_CreateNoBalanceTransaction(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	bnchain := newTestbnchain(api)
	var result string
	err := bnchain.CreateNoBalanceTransaction(&types.NoBalanceTx{TxHex: "0a05636f696e73122c18010a281080c2d72f222131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b7120a08d0630a696c0b3f78dd9ec083a2131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b71"}, &result)
	assert.NoError(t, err)
}

func Testbnchain_ExecWallet(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var testResult interface{}
	in := &rpctypes.ChainExecutor{}
	api.On("ExecWallet", mock.Anything).Return(nil, nil)
	err := client.ExecWallet(in, &testResult)
	assert.NotNil(t, err)
}

func Testbnchain_Query(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var testResult interface{}
	in := rpctypes.Query4Jrpc{Execer: "coins"}
	api.On("Query", mock.Anything).Return(nil, nil)
	err := client.Query(in, &testResult)
	assert.NotNil(t, err)
}

func Testbnchain_DumpPrivkey(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var testResult interface{}
	api.On("DumpPrivkey", mock.Anything).Return(nil, nil)
	err := client.DumpPrivkey(types.ReqString{}, &testResult)
	assert.NoError(t, err)
}

func Testbnchain_GetTotalCoins(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var testResult interface{}
	api.On("StoreGetTotalCoins", mock.Anything).Return(nil, nil)
	err := client.GetTotalCoins(&types.ReqGetTotalCoins{}, &testResult)
	assert.NoError(t, err)
}

func Testbnchain_GetFatalFailure(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var testResult interface{}
	api.On("GetFatalFailure", mock.Anything).Return(&types.Int32{}, nil)
	err := client.GetFatalFailure(nil, &testResult)
	assert.NoError(t, err)
}

func Testbnchain_DecodeRawTransaction(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var testResult interface{}
	//api.On("GetFatalFailure", mock.Anything).Return(&types.Int32{}, nil)
	err := client.DecodeRawTransaction(&types.ReqDecodeRawTransaction{TxHex: "0a05636f696e73122c18010a281080c2d72f222131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b7120a08d0630a696c0b3f78dd9ec083a2131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b71"}, &testResult)
	assert.NoError(t, err)
}

func Testbnchain_CloseQueue(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var testResult interface{}
	api.On("CloseQueue", mock.Anything).Return(nil, nil)
	err := client.CloseQueue(nil, &testResult)
	assert.True(t, testResult.(*types.Reply).IsOk)
	assert.NoError(t, err)
}

func Testbnchain_AddSeqCallBack(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var testResult interface{}
	api.On("AddSeqCallBack", mock.Anything).Return(&types.Reply{}, nil)
	err := client.AddSeqCallBack(nil, &testResult)
	assert.NoError(t, err)
}

func Testbnchain_ListSeqCallBack(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var testResult interface{}
	api.On("ListSeqCallBack", mock.Anything).Return(&types.BlockSeqCBs{}, nil)
	err := client.ListSeqCallBack(nil, &testResult)
	assert.NoError(t, err)
}

func Testbnchain_GetSeqCallBackLastNum(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var testResult interface{}
	api.On("GetSeqCallBackLastNum", mock.Anything).Return(&types.Int64{}, nil)
	err := client.GetSeqCallBackLastNum(nil, &testResult)
	assert.NoError(t, err)
}

func Testbnchain_ConvertExectoAddr(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	client := newTestbnchain(api)
	var testResult string
	err := client.ConvertExectoAddr(rpctypes.ExecNameParm{ExecName: "coins"}, &testResult)
	assert.NoError(t, err)
}

func Test_fmtTxDetail(t *testing.T) {

	tx := &types.Transaction{Execer: []byte("coins")}
	log := &types.ReceiptLog{Ty: 0, Log: []byte("test")}
	receipt := &types.ReceiptData{Ty: 0, Logs: []*types.ReceiptLog{log}}
	detail := &types.TransactionDetail{Tx: tx, Receipt: receipt}
	var err error
	//test withdraw swap from to
	detail.Fromaddr = "from"
	detail.Tx.Payload, err = common.FromHex("0x180322301080c2d72f2205636f696e732a22314761485970576d71414a7371527772706f4e6342385676674b7453776a63487174")
	assert.NoError(t, err)
	tx.To = "to"
	tran, err := fmtTxDetail(detail, false)
	assert.NoError(t, err)
	assert.Equal(t, "to", tran.Fromaddr)
	assert.Equal(t, "from", tx.To)
}

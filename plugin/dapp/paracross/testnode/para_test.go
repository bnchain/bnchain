package testnode

import (
	"testing"

	"github.com/bnchain/bnchain/util"

	_ "github.com/bnchain/bnchain/system"
	_ "github.com/bnchain/plugin/plugin"
)

func TestParaNode(t *testing.T) {
	para := NewParaNode(nil, nil)
	defer para.Close()
	//通过rpc 发生信息
	tx := util.CreateTxWithExecer(para.Para.GetGenesisKey(), "user.p.guodun.none")
	para.Para.SendTxRPC(tx)
	para.Para.WaitHeight(1)
	tx = util.CreateTxWithExecer(para.Para.GetGenesisKey(), "user.p.guodun.none")
	para.Para.SendTxRPC(tx)
	para.Para.WaitHeight(2)
}

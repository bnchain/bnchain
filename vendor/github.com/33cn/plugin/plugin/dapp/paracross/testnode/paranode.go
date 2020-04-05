package testnode

import (
	"github.com/bnchain/bnchain/types"
	"github.com/bnchain/bnchain/util/testnode"
)

/*
1. solo 模式，后台启动一个 主节点
2. 启动一个平行链节点：注意，这个要测试的话，会依赖平行链插件
*/

//ParaNode 平行链节点由两个节点组成
type ParaNode struct {
	Main *testnode.bnchainMock
	Para *testnode.bnchainMock
}

//NewParaNode 创建一个平行链节点
func NewParaNode(main *testnode.bnchainMock, para *testnode.bnchainMock) *ParaNode {
	if main == nil {
		main = testnode.New("", nil)
		main.Listen()
	}
	if para == nil {
		cfg, sub := types.InitCfgString(DefaultConfig)
		testnode.ModifyParaClient(sub, main.GetCfg().RPC.GrpcBindAddr)
		para = testnode.NewWithConfig(cfg, sub, nil)
		para.Listen()
	}
	return &ParaNode{Main: main, Para: para}
}

//Close 关闭系统
func (node *ParaNode) Close() {
	node.Para.Close()
	node.Main.Close()
}

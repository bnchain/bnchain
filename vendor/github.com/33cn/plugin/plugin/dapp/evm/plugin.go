// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package evm

import (
	"github.com/bnchain/bnchain/pluginmgr"
	"github.com/bnchain/plugin/plugin/dapp/evm/commands"
	"github.com/bnchain/plugin/plugin/dapp/evm/executor"
	"github.com/bnchain/plugin/plugin/dapp/evm/rpc"
	"github.com/bnchain/plugin/plugin/dapp/evm/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.ExecutorName,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.EvmCmd,
		RPC:      rpc.Init,
	})
}

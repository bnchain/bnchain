// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unfreeze

import (
	"github.com/bnchain/bnchain/pluginmgr"
	_ "github.com/bnchain/plugin/plugin/dapp/unfreeze/autotest" // register autotest package
	"github.com/bnchain/plugin/plugin/dapp/unfreeze/commands"
	"github.com/bnchain/plugin/plugin/dapp/unfreeze/executor"
	"github.com/bnchain/plugin/plugin/dapp/unfreeze/rpc"
	uf "github.com/bnchain/plugin/plugin/dapp/unfreeze/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     uf.PackageName,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.Cmd,
		RPC:      rpc.Init,
	})
}

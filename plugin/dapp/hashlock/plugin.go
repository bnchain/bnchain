// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashlock

import (
	"github.com/bnchain/bnchain/pluginmgr"
	"github.com/bnchain/plugin/plugin/dapp/hashlock/commands"
	"github.com/bnchain/plugin/plugin/dapp/hashlock/executor"
	"github.com/bnchain/plugin/plugin/dapp/hashlock/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.HashlockX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.HashlockCmd,
		RPC:      nil,
	})
}

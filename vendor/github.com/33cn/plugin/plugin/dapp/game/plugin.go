// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package game

import (
	"github.com/bnchain/bnchain/pluginmgr"
	"github.com/bnchain/plugin/plugin/dapp/game/commands"
	"github.com/bnchain/plugin/plugin/dapp/game/executor"
	gt "github.com/bnchain/plugin/plugin/dapp/game/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     gt.GameX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.Cmd,
		RPC:      nil,
	})
}

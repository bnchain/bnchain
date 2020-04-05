// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package guess

import (
	"github.com/bnchain/bnchain/pluginmgr"
	"github.com/bnchain/plugin/plugin/dapp/guess/commands"
	"github.com/bnchain/plugin/plugin/dapp/guess/executor"
	"github.com/bnchain/plugin/plugin/dapp/guess/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.GuessX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.GuessCmd,
	})
}

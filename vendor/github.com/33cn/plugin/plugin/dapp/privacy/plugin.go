// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package privacy

import (
	"github.com/bnchain/bnchain/pluginmgr"
	_ "github.com/bnchain/plugin/plugin/dapp/privacy/autotest" // register autotest package
	"github.com/bnchain/plugin/plugin/dapp/privacy/commands"
	_ "github.com/bnchain/plugin/plugin/dapp/privacy/crypto" // register crypto package
	"github.com/bnchain/plugin/plugin/dapp/privacy/executor"
	"github.com/bnchain/plugin/plugin/dapp/privacy/rpc"
	"github.com/bnchain/plugin/plugin/dapp/privacy/types"
	_ "github.com/bnchain/plugin/plugin/dapp/privacy/wallet" // register wallet package
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.PrivacyX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.PrivacyCmd,
		RPC:      rpc.Init,
	})
}

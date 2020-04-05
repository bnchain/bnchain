package multisig

import (
	"github.com/bnchain/bnchain/pluginmgr"
	_ "github.com/bnchain/plugin/plugin/dapp/multisig/autotest" //register auto test
	"github.com/bnchain/plugin/plugin/dapp/multisig/commands"
	"github.com/bnchain/plugin/plugin/dapp/multisig/executor"
	"github.com/bnchain/plugin/plugin/dapp/multisig/rpc"
	mty "github.com/bnchain/plugin/plugin/dapp/multisig/types"
	_ "github.com/bnchain/plugin/plugin/dapp/multisig/wallet" // register wallet package
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     mty.MultiSigX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.MultiSigCmd,
		RPC:      rpc.Init,
	})
}

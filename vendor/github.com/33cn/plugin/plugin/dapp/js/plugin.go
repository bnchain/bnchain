package js

import (
	"github.com/bnchain/bnchain/pluginmgr"
	"github.com/bnchain/plugin/plugin/dapp/js/executor"
	ptypes "github.com/bnchain/plugin/plugin/dapp/js/types"

	// init auto test
	_ "github.com/bnchain/plugin/plugin/dapp/js/autotest"
	"github.com/bnchain/plugin/plugin/dapp/js/command"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     ptypes.JsX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      command.JavaScriptCmd,
		RPC:      nil,
	})
}

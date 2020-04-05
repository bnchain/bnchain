// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cert

import (
	"github.com/bnchain/bnchain/pluginmgr"
	"github.com/bnchain/plugin/plugin/dapp/cert/executor"
	"github.com/bnchain/plugin/plugin/dapp/cert/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.CertX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      nil,
		RPC:      nil,
	})
}

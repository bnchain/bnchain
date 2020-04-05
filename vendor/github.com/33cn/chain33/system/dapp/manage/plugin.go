// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package manage manage负责管理配置的插件
// 1. 添加管理
// 2. 添加运营人员
// 3. （未来）修改某些配置项
package manage

import (
	"github.com/bnchain/bnchain/pluginmgr"
	"github.com/bnchain/bnchain/system/dapp/manage/commands"
	"github.com/bnchain/bnchain/system/dapp/manage/executor"
	"github.com/bnchain/bnchain/system/dapp/manage/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.ManageX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.ConfigCmd,
		RPC:      nil,
	})
}

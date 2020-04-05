// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.9

/*
每个系统的功能通过插件完成，插件分成4类：
共识 加密 dapp 存储
这个go 包提供了 官方提供的 插件。
*/
package main

import (
	_ "github.com/bnchain/bnchain/system"
	"github.com/bnchain/bnchain/util/cli"
	_ "github.com/bnchain/plugin/plugin"
	_ "github.com/bnchain/bnchain/plugin"
)

func main() {
	cli.Runbnchain("bnchain")
}

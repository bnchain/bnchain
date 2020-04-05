// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package main

import (
	_ "github.com/bnchain/bnchain/system"
	"github.com/bnchain/bnchain/cli/buildflags"
	_ "github.com/bnchain/bnchain/plugin"

	"github.com/bnchain/bnchain/util/cli"
)

func main() {
	if buildflags.RPCAddr == "" {
		buildflags.RPCAddr = "http://localhost:8801"
	}
	cli.Run(buildflags.RPCAddr, buildflags.ParaName)
}

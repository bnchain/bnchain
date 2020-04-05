// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

// Package main bnchain程序入口
package main

import (
	_ "github.com/bnchain/bnchain/system"
	"github.com/bnchain/bnchain/util/cli"
)

func main() {
	cli.Runbnchain("")
}

// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package system 系统基础插件包
package system

import (
	_ "github.com/bnchain/bnchain/system/consensus/init" //register consensus init package
	_ "github.com/bnchain/bnchain/system/crypto/init"
	_ "github.com/bnchain/bnchain/system/dapp/init"
	_ "github.com/bnchain/bnchain/system/mempool/init"
	_ "github.com/bnchain/bnchain/system/store/init"
)

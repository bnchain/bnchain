/*
 * Copyright Fuzamei Corp. 2018 All Rights Reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package executor

import (
	log "github.com/bnchain/bnchain/common/log/log15"
	drivers "github.com/bnchain/bnchain/system/dapp"
	"github.com/bnchain/bnchain/types"
	oty "github.com/bnchain/plugin/plugin/dapp/oracle/types"
)

var olog = log.New("module", "execs.oracle")
var driverName = oty.OracleX

// Init 执行器初始化
func Init(name string, sub []byte) {
	drivers.Register(newOracle().GetName(), newOracle, types.GetDappFork(driverName, "Enable"))
}

// GetName 获取oracle执行器名
func GetName() string {
	return newOracle().GetName()
}

func newOracle() drivers.Driver {
	t := &oracle{}
	t.SetChild(t)
	t.SetExecutorType(types.LoadExecutorType(driverName))
	return t
}

func init() {
	ety := types.LoadExecutorType(driverName)
	ety.InitFuncList(types.ListMethod(&oracle{}))
}

// oracle driver
type oracle struct {
	drivers.DriverBase
}

func (ora *oracle) GetDriverName() string {
	return oty.OracleX
}

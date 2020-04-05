// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc_test

//only load all plugin and system
import (
	"testing"

	rpctypes "github.com/bnchain/bnchain/rpc/types"
	_ "github.com/bnchain/bnchain/system"
	"github.com/bnchain/bnchain/util/testnode"
	_ "github.com/bnchain/plugin/plugin"
	ty "github.com/bnchain/plugin/plugin/dapp/ticket/types"
	"github.com/stretchr/testify/assert"
)

func TestNewTicket(t *testing.T) {
	cfg, sub := testnode.GetDefaultConfig()
	cfg.Consensus.Name = "ticket"
	mock33 := testnode.NewWithConfig(cfg, sub, nil)
	defer mock33.Close()
	mock33.WaitHeight(5)
	mock33.Listen()
	//选票(可以用hotwallet 关闭选票)
	in := &ty.TicketClose{MinerAddress: mock33.GetHotAddress()}
	var res rpctypes.ReplyHashes
	err := mock33.GetJSONC().Call("ticket.CloseTickets", in, &res)
	assert.Nil(t, err)
}

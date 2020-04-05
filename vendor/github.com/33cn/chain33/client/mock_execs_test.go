// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client_test

import (
	"github.com/bnchain/bnchain/queue"
	"github.com/bnchain/bnchain/types"
)

type mockExecs struct {
}

func (m *mockExecs) SetQueueClient(q queue.Queue) {
	go func() {
		topic := "execs"
		client := q.Client()
		client.Sub(topic)
		for msg := range client.Recv() {
			switch msg.Ty {
			case types.EventBlockChainQuery:
				msg.Reply(client.NewMessage(topic, types.EventBlockChainQuery, &types.Reply{}))
			default:
				msg.ReplyErr("Do not support", types.ErrNotSupport)
			}
		}
	}()
}

func (m *mockExecs) Close() {
}

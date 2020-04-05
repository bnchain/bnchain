// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package timeline

import (
	"encoding/json"
	"testing"

	"github.com/bnchain/bnchain/system/mempool"
	"github.com/bnchain/bnchain/types"
)

func TestNewMempool(t *testing.T) {
	sub, _ := json.Marshal(&mempool.SubConfig{PoolCacheSize: 2})
	module := New(&types.Mempool{}, sub)
	mem := module.(*mempool.Mempool)
	mem.Close()
}

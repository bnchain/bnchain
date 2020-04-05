// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package store store the world - state data
package store

import (
	"github.com/bnchain/bnchain/queue"
	"github.com/bnchain/bnchain/system/store"
	"github.com/bnchain/bnchain/types"
)

// New new store queue module
func New(cfg *types.Store, sub map[string][]byte) queue.Module {
	s, err := store.Load(cfg.Name)
	if err != nil {
		panic("Unsupported store type:" + cfg.Name + " " + err.Error())
	}
	subcfg, ok := sub[cfg.Name]
	if !ok {
		subcfg = nil
	}
	return s(cfg, subcfg)
}

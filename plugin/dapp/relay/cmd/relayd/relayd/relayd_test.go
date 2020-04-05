// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package relayd

import (
	"testing"

	"github.com/bnchain/bnchain/common"
	"github.com/bnchain/bnchain/common/address"
	"github.com/bnchain/bnchain/common/crypto"
	"github.com/bnchain/bnchain/types"
)

func TestGeneratePrivateKey(t *testing.T) {
	cr, err := crypto.New(types.GetSignName("", types.SECP256K1))
	if err != nil {
		t.Fatal(err)
	}

	key, err := cr.GenKey()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("private key: ", common.ToHex(key.Bytes()))
	t.Log("publick key: ", common.ToHex(key.PubKey().Bytes()))
	t.Log("    address: ", address.PubKeyToAddress(key.PubKey().Bytes()))
}

package executor

import (
	"github.com/bnchain/bnchain/types"
	ptypes "github.com/bnchain/plugin/plugin/dapp/js/types"
)

func calcAllPrefix(name string) ([]byte, []byte) {
	execer := types.ExecName("user." + ptypes.JsX + "." + name)
	state := types.CalcStatePrefix([]byte(execer))
	local := types.CalcLocalPrefix([]byte(execer))
	return state, local
}

func calcCodeKey(name string) []byte {
	return append([]byte("mavl-"+ptypes.JsX+"-code-"), []byte(name)...)
}

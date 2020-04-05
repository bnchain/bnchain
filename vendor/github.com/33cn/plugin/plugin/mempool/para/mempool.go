package para

import (
	"github.com/bnchain/bnchain/queue"
	drivers "github.com/bnchain/bnchain/system/mempool"
	"github.com/bnchain/bnchain/types"
)

//--------------------------------------------------------------------------------
// Module Mempool

func init() {
	drivers.Reg("para", New)
}

//New 创建price cache 结构的 mempool
func New(cfg *types.Mempool, sub []byte) queue.Module {
	return NewMempool(cfg)
}

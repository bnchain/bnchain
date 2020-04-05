package init

import (
	_ "github.com/bnchain/plugin/plugin/consensus/para"       //auto gen
	_ "github.com/bnchain/plugin/plugin/consensus/pbft"       //auto gen
	_ "github.com/bnchain/plugin/plugin/consensus/raft"       //auto gen
	_ "github.com/bnchain/plugin/plugin/consensus/tendermint" //auto gen
	_ "github.com/bnchain/plugin/plugin/consensus/ticket"     //auto gen
)

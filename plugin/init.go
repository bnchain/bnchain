package plugin

import (
	_ "github.com/bnchain/plugin/plugin/consensus/init" //consensus init
	_ "github.com/bnchain/plugin/plugin/crypto/init"    //crypto init
	_ "github.com/bnchain/plugin/plugin/dapp/init"      //dapp init
	_ "github.com/bnchain/plugin/plugin/mempool/init"   //mempool init
	_ "github.com/bnchain/plugin/plugin/store/init"     //store init
)

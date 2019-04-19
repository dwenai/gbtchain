package plugin

import (
	_ "github.com/dwenai/gbtchain/plugin/consensus/init" //consensus init
	_ "github.com/dwenai/gbtchain/plugin/crypto/init"    //crypto init
	_ "github.com/dwenai/gbtchain/plugin/dapp/init"      //dapp init
	_ "github.com/dwenai/gbtchain/plugin/mempool/init"   //mempool init
	_ "github.com/dwenai/gbtchain/plugin/store/init"     //store init
)

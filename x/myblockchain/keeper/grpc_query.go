package keeper

import (
	"github.com/cosmostest/myblockchain/x/myblockchain/types"
)

var _ types.QueryServer = Keeper{}

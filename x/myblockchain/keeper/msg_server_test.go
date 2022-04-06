package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/cosmostest/myblockchain/testutil/keeper"
	"github.com/cosmostest/myblockchain/x/myblockchain/keeper"
	"github.com/cosmostest/myblockchain/x/myblockchain/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MyblockchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Zireael26/checkers/testutil/keeper"
	"github.com/Zireael26/checkers/x/checkers"
	"github.com/Zireael26/checkers/x/checkers/keeper"
	"github.com/Zireael26/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func setupMsgServerCreateGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreateGame(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "1",
	}, *createResponse)
}

func TestCreate1GameHasSaved(t *testing.T) {
	msgServer, keeper, context := setupMsgServerCreateGame(t)
	msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	systemInfo, found := keeper.GetSystemInfo(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId: 2,
	}, systemInfo)

	game1, found1 := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found1)
	require.EqualValues(t, types.StoredGame{
		Index: "1",
		Board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:  "b",
		Black: bob,
		Red:   carol,
	}, game1)
}

func TestCreate3Games(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)
	createGameResponse1, err1 := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.Nil(t, err1)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "1",
	}, *createGameResponse1)

	createGameResponse2, err2 := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
	})
	require.Nil(t, err2)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "2",
	}, *createGameResponse2)

	createGameResponse3, err3 := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Black:   alice,
		Red:     bob,
	})
	require.Nil(t, err3)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "3",
	}, *createGameResponse3)
}

func TestCreate3GamesHasSaved(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
	})
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Black:   alice,
		Red:     bob,
	})
	systemInfo, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId: 4,
	}, systemInfo)
	game1, found1 := keeper.GetStoredGame(ctx, "1")
	require.True(t, found1)
	require.EqualValues(t, types.StoredGame{
		Index: "1",
		Board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:  "b",
		Black: bob,
		Red:   carol,
	}, game1)
	game2, found2 := keeper.GetStoredGame(ctx, "2")
	require.True(t, found2)
	require.EqualValues(t, types.StoredGame{
		Index: "2",
		Board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:  "b",
		Black: carol,
		Red:   alice,
	}, game2)
	game3, found3 := keeper.GetStoredGame(ctx, "3")
	require.True(t, found3)
	require.EqualValues(t, types.StoredGame{
		Index: "3",
		Board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:  "b",
		Black: alice,
		Red:   bob,
	}, game3)
}

func TestCreate3GamesGetAll(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
	})
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Black:   alice,
		Red:     bob,
	})
	games := keeper.GetAllStoredGame(sdk.UnwrapSDKContext(context))
	require.Len(t, games, 3)
	require.EqualValues(t, types.StoredGame{
		Index: "1",
		Board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:  "b",
		Black: bob,
		Red:   carol,
	}, games[0])
	require.EqualValues(t, types.StoredGame{
		Index: "2",
		Board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:  "b",
		Black: carol,
		Red:   alice,
	}, games[1])
	require.EqualValues(t, types.StoredGame{
		Index: "3",
		Board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:  "b",
		Black: alice,
		Red:   bob,
	}, games[2])
}

func TestCreateGameFarFuture(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	systemInfo, found := keeper.GetSystemInfo(ctx)
	systemInfo.NextId = 1024
	keeper.SetSystemInfo(ctx, systemInfo)
	createResponse, err := msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "1024",
	}, *createResponse)
	systemInfo, found = keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId: 1025,
	}, systemInfo)
	game1, found1 := keeper.GetStoredGame(ctx, "1024")
	require.True(t, found1)
	require.EqualValues(t, types.StoredGame{
		Index: "1024",
		Board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:  "b",
		Black: bob,
		Red:   carol,
	}, game1)
}
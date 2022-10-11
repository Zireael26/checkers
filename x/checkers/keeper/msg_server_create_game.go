package keeper

import (
	"context"
	"strconv"

	"github.com/Zireael26/checkers/x/checkers/rules"
	"github.com/Zireael26/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get SystemInfo to get the ID for storing the game
	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)

	// Create the game
	newGame := rules.New()
	storedGame := types.StoredGame{
		Index:       newIndex,
		Board:       newGame.String(),
		Red:         msg.Red,
		Black:       msg.Black,
		Turn:        rules.PieceStrings[newGame.Turn],
		MoveCount:   0,
		BeforeIndex: types.NoFifoIndex,
		AfterIndex:  types.NoFifoIndex,
		Deadline:    types.FormatDeadline(types.GetNextDeadline(ctx)),
		Winner:      rules.PieceStrings[rules.NO_PLAYER],
	}

	// Validate the created game
	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}

	// Set game to FIFO Tail
	k.Keeper.SendToFifoTail(ctx, &storedGame, &systemInfo)

	// Store the game
	k.Keeper.SetStoredGame(ctx, storedGame)

	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.GameCreatedEventType,
			sdk.NewAttribute(types.GameCreatedEventCreator, msg.Creator),
			sdk.NewAttribute(types.GameCreatedEventGameIndex, newIndex),
			sdk.NewAttribute(types.GameCreatedEventBlack, msg.Black),
			sdk.NewAttribute(types.GameCreatedEventRed, msg.Red),
		),
	)

	return &types.MsgCreateGameResponse{GameIndex: newIndex}, nil
}

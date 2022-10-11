package keeper

import (
	"context"
	"strconv"

	"github.com/Zireael26/checkers/x/checkers/rules"
	"github.com/Zireael26/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) PlayMove(goCtx context.Context, msg *types.MsgPlayMove) (*types.MsgPlayMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedGame, found := k.Keeper.GetStoredGame(ctx, msg.GameIndex)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "%s", msg.GameIndex)
	}

	if storedGame.Winner != rules.PieceStrings[rules.NO_PLAYER] {
		return nil, types.ErrGameFinished
	}

	isBlack := storedGame.Black == msg.Creator
	isRed := storedGame.Red == msg.Creator
	var player rules.Player
	if !isBlack && !isRed {
		return nil, sdkerrors.Wrapf(types.ErrCreatorNotPlayer, "%s", msg.Creator)
	} else if isBlack && isRed {
		player = rules.StringPieces[storedGame.Turn].Player
	} else if isBlack {
		player = rules.BLACK_PLAYER
	} else {
		player = rules.RED_PLAYER
	}

	game, parseErr := storedGame.Parse()
	if parseErr != nil {
		panic(parseErr.Error())
	}

	if !game.TurnIs(player) {
		return nil, sdkerrors.Wrapf(types.ErrNotPlayerTurn, "%s", player)
	}

	captured, moveErr := game.Move(
		rules.Pos{X: int(msg.FromX), Y: int(msg.FromY)},
		rules.Pos{X: int(msg.ToX), Y: int(msg.ToY)},
	)
	if moveErr != nil {
		return nil, sdkerrors.Wrapf(types.ErrWrongMove, moveErr.Error())
	}

	// Get SystemInfo for FIFO
	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	// Set game to FIFO Tail
	k.Keeper.SendToFifoTail(ctx, &storedGame, &systemInfo)

	storedGame.MoveCount++
	storedGame.Deadline = types.FormatDeadline(types.GetNextDeadline(ctx))

	// Store SystemInfo with new FIFO info
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	storedGame.Board = game.String()
	storedGame.Turn = rules.PieceStrings[game.Turn]

	storedGame.Winner = rules.PieceStrings[game.Winner()]

	lastBoard := game.String()
	if storedGame.Winner == rules.PieceStrings[rules.NO_PLAYER] {
		k.Keeper.SendToFifoTail(ctx, &storedGame, &systemInfo)
		storedGame.Board = lastBoard
	} else {
		k.Keeper.RemoveFromFifo(ctx, &storedGame, &systemInfo)
		storedGame.Board = ""
	}
	k.Keeper.SetStoredGame(ctx, storedGame)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.MovePlayedEventType,
			sdk.NewAttribute(types.MovePlayedEventCreator, msg.Creator),
			sdk.NewAttribute(types.MovePlayedEventGameIndex, msg.GameIndex),
			sdk.NewAttribute(types.MovePlayedEventCapturedX, strconv.FormatInt(int64(captured.X), 10)),
			sdk.NewAttribute(types.MovePlayedEventCapturedY, strconv.FormatInt(int64(captured.Y), 10)),
			sdk.NewAttribute(types.MovePlayedEventWinner, rules.PieceStrings[game.Winner()]),
			sdk.NewAttribute(types.MovePlayedEventBoard, lastBoard),
		),
	)

	return &types.MsgPlayMoveResponse{
		CapturedX: int32(captured.X),
		CapturedY: int32(captured.Y),
		Winner:    rules.PieceStrings[game.Winner()],
	}, nil
}

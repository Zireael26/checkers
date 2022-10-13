package keeper

import (
	"context"
	"fmt"

	"github.com/Zireael26/checkers/x/checkers/rules"
	"github.com/Zireael26/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ForfeitExpiredGames(goCtx context.Context) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	opponents := map[string]string{
		rules.PieceStrings[rules.BLACK_PLAYER]: rules.PieceStrings[rules.RED_PLAYER],
		rules.PieceStrings[rules.RED_PLAYER]:   rules.PieceStrings[rules.BLACK_PLAYER],
	}

	systemInfo, found := k.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}

	gameIndex := systemInfo.FifoHeadIndex
	var storedGame types.StoredGame

	for {
		if gameIndex == types.NoFifoIndex {
			break
		}

		storedGame, found = k.GetStoredGame(ctx, gameIndex)
		if !found {
			panic("Fifo head game not found: " + systemInfo.FifoHeadIndex)
		}
		deadline, err := storedGame.GetDeadlineAsTime()
		if err != nil {
			panic(err)
		}

		if deadline.Before(ctx.BlockTime()) {
			// Remove expired game from FIFO
			k.RemoveFromFifo(ctx, &storedGame, &systemInfo)

			lastBoard := storedGame.Board
			if storedGame.MoveCount <= 1 {
				// No point in keeping a game that was never really played
				k.RemoveStoredGame(ctx, gameIndex)
			} else {
				storedGame.Winner, found = opponents[storedGame.Turn]
				if !found {
					panic(fmt.Sprintf(types.ErrCannotFindWinnerByColor.Error(), storedGame.Turn))
				}
				storedGame.Board = ""
				k.SetStoredGame(ctx, storedGame)
			}

			// Set winner
			storedGame.Winner = opponents[storedGame.Turn]
			k.SetStoredGame(ctx, storedGame)

			ctx.EventManager().EmitEvent(
				sdk.NewEvent(types.GameForfeitedEventType,
					sdk.NewAttribute(types.GameForfeitedEventGameIndex, storedGame.Index),
					sdk.NewAttribute(types.GameForfeitedEventWinner, storedGame.Winner),
					sdk.NewAttribute(types.GameForfeitedEventBoard, lastBoard),
				),
			)
		} else {
			break
		}

		gameIndex = storedGame.AfterIndex
	}

	k.SetSystemInfo(ctx, systemInfo)

}

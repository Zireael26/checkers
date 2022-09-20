package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidBlack     = sdkerrors.Register(ModuleName, 1100, "invalid black address: (%s)")
	ErrInvalidRed       = sdkerrors.Register(ModuleName, 1101, "invalid red address: (%s)")
	ErrGameNotParseable = sdkerrors.Register(ModuleName, 1102, "game cannot be parsed")
	ErrGameNotFound     = sdkerrors.Register(ModuleName, 1103, "game by id not found")
	ErrCreatorNotPlayer = sdkerrors.Register(ModuleName, 1104, "message creator is not a player")
	ErrNotPlayerTurn    = sdkerrors.Register(ModuleName, 1105, "player tried to play out of turn")
	ErrWrongMove        = sdkerrors.Register(ModuleName, 1106, "wrong move")
)

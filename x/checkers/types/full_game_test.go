package types_test

import (
	"strings"
	"testing"

	rules "github.com/Zireael26/checkers/x/checkers/rules"
	"github.com/Zireael26/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func GetStoredGame1() *types.StoredGame {
	return &types.StoredGame{
		Black: alice,
		Red:   bob,
		Index: "1",
		Board: rules.New().String(),
		Turn:  "b",
	}
}

func TestCanGetAddressBlack(t *testing.T) {
	aliceAddress, err1 := sdk.AccAddressFromBech32(alice)
	black, err2 := GetStoredGame1().GetBlackAddress()
	require.Equal(t, aliceAddress, black)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongBlack(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Black = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4"
	black, err := storedGame.GetBlackAddress()
	require.Nil(t, black)
	require.EqualError(t, err, "invalid black address: (cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4): decoding bech32 failed: invalid checksum (expected 3xn9d3 got 3xn9d4)")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestCanGetAddressRed(t *testing.T) {
	bobAddress, err1 := sdk.AccAddressFromBech32(bob)
	red, err2 := GetStoredGame1().GetRedAddress()
	require.Equal(t, bobAddress, red)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongRed(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Red = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8h"
	red, err := storedGame.GetRedAddress()
	require.Nil(t, red)
	require.EqualError(t, err, "invalid red address: (cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8h): decoding bech32 failed: invalid checksum (expected xqhc8g got xqhc8h)")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestParseGameCorrect(t *testing.T) {
	game, err := GetStoredGame1().Parse()
	require.EqualValues(t, rules.New().Pieces, game.Pieces)
	require.Nil(t, err)
}

func TestParseGameIfChangedOk(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Board = strings.Replace(storedGame.Board, "b", "r", 1)
	game, err := storedGame.Parse()
	require.NotEqualValues(t, rules.New().Pieces, game.Pieces)
	require.Nil(t, err)
}

func TestParseGameWrongPieceColor(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Board = strings.Replace(storedGame.Board, "b", "w", 1)
	game, err := storedGame.Parse()
	require.Nil(t, game)
	require.EqualError(t, err, "game cannot be parsed: invalid board, invalid piece at 1, 0")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestParseGameWrongTurnColor(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Turn = "w"
	game, err := storedGame.Parse()
	require.Nil(t, game)
	require.EqualError(t, err, "game cannot be parsed: Turn: w")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestGameValidateOk(t *testing.T) {
	storedGame := GetStoredGame1()
	require.NoError(t, storedGame.Validate())
}

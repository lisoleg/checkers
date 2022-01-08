package keeper

import (
	"context"
	"strings"

	"github.com/alice/checkers/x/checkers/rules"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) PlayMove(goCtx context.Context, msg *types.MsgPlayMove) (*types.MsgPlayMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedGame, found := k.Keeper.GetStoredGame(ctx, msg.IdValue)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "game not found %s", msg.IdValue)
	}

	var player rules.Player
	if strings.Compare(storedGame.Red, msg.Creator) == 0 {
		player = rules.RED_PLAYER
	} else if strings.Compare(storedGame.Black, msg.Creator) == 0 {
		player = rules.BLACK_PLAYER
	} else {
		return nil, types.ErrCreatorNotPlayer
	}

	game, err := storedGame.ParseGame()
	if err != nil {
		panic(err.Error())
	}

	if !game.TurnIs(player) {
    return nil, types.ErrNotPlayerTurn
	}

	captured, moveErr := game.Move(
    rules.Pos{
			X: int(msg.FromX),
			Y: int(msg.FromY),
    },
    rules.Pos{
			X: int(msg.ToX),
			Y: int(msg.ToY),
    },
	)
	if moveErr != nil {
		return nil, sdkerrors.Wrapf(moveErr, types.ErrWrongMove.Error())
	}

	storedGame.Game = game.String()
	storedGame.Turn = game.Turn.Color
	k.Keeper.SetStoredGame(ctx, storedGame)

	return &types.MsgPlayMoveResponse{
    IdValue:   msg.IdValue,
    CapturedX: int64(captured.X),
    CapturedY: int64(captured.Y),
    Winner:    game.Winner().Color,
}, nil
}

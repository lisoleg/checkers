package keeper

import (
	"fmt"

	"github.com/alice/checkers/x/checkers/rules"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k *Keeper) MustAddWonGameResultToPlayer(ctx sdk.Context, player sdk.AccAddress) types.PlayerInfo {
	return mustAddDeltaGameResultToPlayer(k, ctx, player, 1, 0, 0)
}

func (k *Keeper) MustAddLostGameResultToPlayer(ctx sdk.Context, player sdk.AccAddress) types.PlayerInfo {
	return mustAddDeltaGameResultToPlayer(k, ctx, player, 0, 1, 0)
}

func (k *Keeper) MustAddForfeitedGameResultToPlayer(ctx sdk.Context, player sdk.AccAddress) types.PlayerInfo {
	return mustAddDeltaGameResultToPlayer(k, ctx, player, 0, 0, 1)
}

func (k *Keeper) MustRegisterPlayerWin(ctx sdk.Context, storedGame *types.StoredGame) (winnerInfo types.PlayerInfo, loserInfo types.PlayerInfo) {
	winnerAddress, loserAddress := getWinnerAndLoserAddresses(storedGame)
	return k.MustAddWonGameResultToPlayer(ctx, winnerAddress),
		k.MustAddLostGameResultToPlayer(ctx, loserAddress)
}

func (k *Keeper) MustRegisterPlayerForfeit(ctx sdk.Context, storedGame *types.StoredGame) (winnerInfo types.PlayerInfo, forfeiterInfo types.PlayerInfo) {
	winnerAddress, loserAddress := getWinnerAndLoserAddresses(storedGame)
	return k.MustAddWonGameResultToPlayer(ctx, winnerAddress),
		k.MustAddForfeitedGameResultToPlayer(ctx, loserAddress)
}

func getWinnerAndLoserAddresses(storedGame *types.StoredGame) (winnerAddress sdk.AccAddress, loserAddress sdk.AccAddress) {
	if storedGame.Winner == rules.NO_PLAYER.Color {
		panic(types.ErrThereIsNoWinner.Error())
	}
	redAddress, err := storedGame.GetRedAddress()
	if err != nil {
		panic(err.Error())
	}
	blackAddress, err := storedGame.GetBlackAddress()
	if err != nil {
		panic(err.Error())
	}
	if storedGame.Winner == rules.RED_PLAYER.Color {
		winnerAddress = redAddress
		loserAddress = blackAddress
	} else if storedGame.Winner == rules.BLACK_PLAYER.Color {
		winnerAddress = blackAddress
		loserAddress = redAddress
	} else {
		panic(fmt.Sprintf(types.ErrWinnerNotParseable.Error(), storedGame.Winner))
	}
	return winnerAddress, loserAddress
}

func mustAddDeltaGameResultToPlayer(
	k *Keeper,
	ctx sdk.Context,
	player sdk.AccAddress,
	wonDelta uint64,
	lostDelta uint64,
	forfeitedDelta uint64,
) (playerInfo types.PlayerInfo) {
	playerInfo, found := k.GetPlayerInfo(ctx, player.String())
	if !found {
		playerInfo = types.PlayerInfo{
			Index:          player.String(),
			WonCount:       0,
			LostCount:      0,
			ForfeitedCount: 0,
		}
	}
	playerInfo.WonCount += wonDelta
	playerInfo.LostCount += lostDelta
	playerInfo.ForfeitedCount += forfeitedDelta
	k.SetPlayerInfo(ctx, playerInfo)
	return playerInfo
}

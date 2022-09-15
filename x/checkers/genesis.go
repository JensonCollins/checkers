package checkers

import (
	"github.com/JensonCollins/checkers/x/checkers/keeper"
	"github.com/JensonCollins/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.NextGame != nil {
		k.SetNextGame(ctx, *genState.NextGame)
	}
	// Set all the storedGame
	for _, elem := range genState.StoredGameList {
		k.SetStoredGame(ctx, elem)
	}
	// Set if defined
	if genState.SystemInfo != nil {
		k.SetSystemInfo(ctx, *genState.SystemInfo)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all nextGame
	nextGame, found := k.GetNextGame(ctx)
	if found {
		genesis.NextGame = &nextGame
	}
	genesis.StoredGameList = k.GetAllStoredGame(ctx)
	// Get all systemInfo
	systemInfo, found := k.GetSystemInfo(ctx)
	if found {
		genesis.SystemInfo = &systemInfo
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

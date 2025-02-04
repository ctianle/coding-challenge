package crudechain_test

import (
	"testing"

	keepertest "crudechain/testutil/keeper"
	"crudechain/testutil/nullify"
	crudechain "crudechain/x/crudechain/module"
	"crudechain/x/crudechain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CrudeblockList: []types.Crudeblock{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CrudeblockCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CrudechainKeeper(t)
	crudechain.InitGenesis(ctx, k, genesisState)
	got := crudechain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CrudeblockList, got.CrudeblockList)
	require.Equal(t, genesisState.CrudeblockCount, got.CrudeblockCount)
	// this line is used by starport scaffolding # genesis/test/assert
}

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

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CrudechainKeeper(t)
	crudechain.InitGenesis(ctx, k, genesisState)
	got := crudechain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "crudchain/testutil/keeper"
	"crudchain/x/crudchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CrudchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

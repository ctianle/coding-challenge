package keeper_test

import (
	"context"
	"testing"

	keepertest "crudechain/testutil/keeper"
	"crudechain/testutil/nullify"
	"crudechain/x/crudechain/keeper"
	"crudechain/x/crudechain/types"

	"github.com/stretchr/testify/require"
)

func createNCrudeblock(keeper keeper.Keeper, ctx context.Context, n int) []types.Crudeblock {
	items := make([]types.Crudeblock, n)
	for i := range items {
		items[i].Id = keeper.AppendCrudeblock(ctx, items[i])
	}
	return items
}

func TestCrudeblockGet(t *testing.T) {
	keeper, ctx := keepertest.CrudechainKeeper(t)
	items := createNCrudeblock(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetCrudeblock(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestCrudeblockRemove(t *testing.T) {
	keeper, ctx := keepertest.CrudechainKeeper(t)
	items := createNCrudeblock(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCrudeblock(ctx, item.Id)
		_, found := keeper.GetCrudeblock(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCrudeblockGetAll(t *testing.T) {
	keeper, ctx := keepertest.CrudechainKeeper(t)
	items := createNCrudeblock(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCrudeblock(ctx)),
	)
}

func TestCrudeblockCount(t *testing.T) {
	keeper, ctx := keepertest.CrudechainKeeper(t)
	items := createNCrudeblock(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCrudeblockCount(ctx))
}

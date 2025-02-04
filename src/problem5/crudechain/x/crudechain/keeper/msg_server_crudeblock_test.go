package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"crudechain/x/crudechain/types"
)

func TestCrudeblockMsgServerCreate(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateCrudeblock(wctx, &types.MsgCreateCrudeblock{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestCrudeblockMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateCrudeblock
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateCrudeblock{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateCrudeblock{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateCrudeblock{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateCrudeblock(wctx, &types.MsgCreateCrudeblock{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateCrudeblock(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCrudeblockMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteCrudeblock
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteCrudeblock{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteCrudeblock{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteCrudeblock{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateCrudeblock(wctx, &types.MsgCreateCrudeblock{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteCrudeblock(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

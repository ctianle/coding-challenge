package keeper

import (
	"context"

	"crudechain/x/crudechain/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CrudeblockAll(ctx context.Context, req *types.QueryAllCrudeblockRequest) (*types.QueryAllCrudeblockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var crudeblocks []types.Crudeblock

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	crudeblockStore := prefix.NewStore(store, types.KeyPrefix(types.CrudeblockKey))

	pageRes, err := query.Paginate(crudeblockStore, req.Pagination, func(key []byte, value []byte) error {
		var crudeblock types.Crudeblock
		if err := k.cdc.Unmarshal(value, &crudeblock); err != nil {
			return err
		}

		crudeblocks = append(crudeblocks, crudeblock)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCrudeblockResponse{Crudeblock: crudeblocks, Pagination: pageRes}, nil
}

func (k Keeper) Crudeblock(ctx context.Context, req *types.QueryGetCrudeblockRequest) (*types.QueryGetCrudeblockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	crudeblock, found := k.GetCrudeblock(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetCrudeblockResponse{Crudeblock: crudeblock}, nil
}

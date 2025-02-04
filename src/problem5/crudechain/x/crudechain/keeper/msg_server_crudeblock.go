package keeper

import (
	"context"
	"fmt"

	"crudechain/x/crudechain/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateCrudeblock(goCtx context.Context, msg *types.MsgCreateCrudeblock) (*types.MsgCreateCrudeblockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var crudeblock = types.Crudeblock{
		Creator:     msg.Creator,
		Name:        msg.Name,
		Description: msg.Description,
	}

	id := k.AppendCrudeblock(
		ctx,
		crudeblock,
	)

	return &types.MsgCreateCrudeblockResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateCrudeblock(goCtx context.Context, msg *types.MsgUpdateCrudeblock) (*types.MsgUpdateCrudeblockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var crudeblock = types.Crudeblock{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Name:        msg.Name,
		Description: msg.Description,
	}

	// Checks that the element exists
	val, found := k.GetCrudeblock(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetCrudeblock(ctx, crudeblock)

	return &types.MsgUpdateCrudeblockResponse{}, nil
}

func (k msgServer) DeleteCrudeblock(goCtx context.Context, msg *types.MsgDeleteCrudeblock) (*types.MsgDeleteCrudeblockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetCrudeblock(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveCrudeblock(ctx, msg.Id)

	return &types.MsgDeleteCrudeblockResponse{}, nil
}

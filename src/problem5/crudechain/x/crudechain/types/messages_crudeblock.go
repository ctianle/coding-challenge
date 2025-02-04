package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateCrudeblock{}

func NewMsgCreateCrudeblock(creator string, name string, description string) *MsgCreateCrudeblock {
	return &MsgCreateCrudeblock{
		Creator:     creator,
		Name:        name,
		Description: description,
	}
}

func (msg *MsgCreateCrudeblock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCrudeblock{}

func NewMsgUpdateCrudeblock(creator string, id uint64, name string, description string) *MsgUpdateCrudeblock {
	return &MsgUpdateCrudeblock{
		Id:          id,
		Creator:     creator,
		Name:        name,
		Description: description,
	}
}

func (msg *MsgUpdateCrudeblock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteCrudeblock{}

func NewMsgDeleteCrudeblock(creator string, id uint64) *MsgDeleteCrudeblock {
	return &MsgDeleteCrudeblock{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteCrudeblock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

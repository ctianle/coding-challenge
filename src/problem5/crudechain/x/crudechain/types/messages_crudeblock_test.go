package types

import (
	"testing"

	"crudechain/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateCrudeblock_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateCrudeblock
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateCrudeblock{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateCrudeblock{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateCrudeblock_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateCrudeblock
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateCrudeblock{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateCrudeblock{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteCrudeblock_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteCrudeblock
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteCrudeblock{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteCrudeblock{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

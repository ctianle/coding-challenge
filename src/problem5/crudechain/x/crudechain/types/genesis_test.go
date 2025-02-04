package types_test

import (
	"testing"

	"crudechain/x/crudechain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				CrudeblockList: []types.Crudeblock{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				CrudeblockCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated crudeblock",
			genState: &types.GenesisState{
				CrudeblockList: []types.Crudeblock{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid crudeblock count",
			genState: &types.GenesisState{
				CrudeblockList: []types.Crudeblock{
					{
						Id: 1,
					},
				},
				CrudeblockCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

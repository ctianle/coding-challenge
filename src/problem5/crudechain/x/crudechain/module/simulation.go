package crudechain

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"crudechain/testutil/sample"
	crudechainsimulation "crudechain/x/crudechain/simulation"
	"crudechain/x/crudechain/types"
)

// avoid unused import issue
var (
	_ = crudechainsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateCrudeblock = "op_weight_msg_crudeblock"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCrudeblock int = 100

	opWeightMsgUpdateCrudeblock = "op_weight_msg_crudeblock"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCrudeblock int = 100

	opWeightMsgDeleteCrudeblock = "op_weight_msg_crudeblock"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteCrudeblock int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	crudechainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		CrudeblockList: []types.Crudeblock{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		CrudeblockCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&crudechainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateCrudeblock int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateCrudeblock, &weightMsgCreateCrudeblock, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCrudeblock = defaultWeightMsgCreateCrudeblock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCrudeblock,
		crudechainsimulation.SimulateMsgCreateCrudeblock(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCrudeblock int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateCrudeblock, &weightMsgUpdateCrudeblock, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCrudeblock = defaultWeightMsgUpdateCrudeblock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCrudeblock,
		crudechainsimulation.SimulateMsgUpdateCrudeblock(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteCrudeblock int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteCrudeblock, &weightMsgDeleteCrudeblock, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCrudeblock = defaultWeightMsgDeleteCrudeblock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteCrudeblock,
		crudechainsimulation.SimulateMsgDeleteCrudeblock(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateCrudeblock,
			defaultWeightMsgCreateCrudeblock,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crudechainsimulation.SimulateMsgCreateCrudeblock(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateCrudeblock,
			defaultWeightMsgUpdateCrudeblock,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crudechainsimulation.SimulateMsgUpdateCrudeblock(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteCrudeblock,
			defaultWeightMsgDeleteCrudeblock,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crudechainsimulation.SimulateMsgDeleteCrudeblock(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

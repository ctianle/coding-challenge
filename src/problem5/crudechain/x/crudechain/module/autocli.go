package crudechain

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "crudechain/api/crudechain/crudechain"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "CrudeblockAll",
					Use:       "list-crudeblock",
					Short:     "List all crudeblock",
				},
				{
					RpcMethod:      "Crudeblock",
					Use:            "show-crudeblock [id]",
					Short:          "Shows a crudeblock by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateCrudeblock",
					Use:            "create-crudeblock [name] [description]",
					Short:          "Create crudeblock",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "description"}},
				},
				{
					RpcMethod:      "UpdateCrudeblock",
					Use:            "update-crudeblock [id] [name] [description]",
					Short:          "Update crudeblock",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "name"}, {ProtoField: "description"}},
				},
				{
					RpcMethod:      "DeleteCrudeblock",
					Use:            "delete-crudeblock [id]",
					Short:          "Delete crudeblock",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

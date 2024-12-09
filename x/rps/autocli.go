package rps

import (
	rpsv1 "challenge/x/rps/types"
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: nil,
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: rpsv1.Msg_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "CreateStudent",
					Use:       "create student",
					Short:     "creates a new student",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
						{ProtoField: "age"},
					},
				},
			},
		},
	}
}

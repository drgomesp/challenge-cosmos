package rps

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Tx:    &autocliv1.ServiceCommandDescriptor{},
		Query: &autocliv1.ServiceCommandDescriptor{},
	}
}

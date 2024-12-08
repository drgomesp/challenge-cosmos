package rps

import (
	"encoding/json"
	"fmt"

	rpsKeeper "challenge/x/rps/keeper"
	"challenge/x/rps/types"

	"cosmossdk.io/core/appmodule"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
)

var (
	// _ module.AppModuleBasic = AppModule{} // Module Name and codification settings
	_ module.HasGenesis   = AppModule{} // Handlers for the genesis handle (export and initial state)
	_ appmodule.AppModule = AppModule{} // Register services*, genesis handle and how to interact with consensours mecanism
)

// ConsensusVersion defines the current module consensus version.
const ConsensusVersion = 1

type AppModule struct {
	cdc    codec.Codec      // Responsible to handle the codification and decodification
	keeper rpsKeeper.Keeper // Responsible for keeper handle (keepes handle state and module's logic)
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper rpsKeeper.Keeper) AppModule {
	return AppModule{
		cdc:    cdc,
		keeper: keeper,
	}
}

// ********************* IMPLEMENT AppModuleBasic INTERFACE ******************
// Name returns the checkers module's name.
func (AppModule) Name() string { return types.ModuleName }

// RegisterLegacyAminoCodec registers the checkers module's types on the LegacyAmino codec.
// New modules do not need to support Amino.
func (AppModule) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the example module.
// func (AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
// 	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
// 		panic(err)
// 	}
// }

// RegisterInterfaces registers interfaces and implementations of the checkers module.
func (AppModule) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// ****************************************************************************

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries.
// func (am AppModule) RegisterServices(cfg module.Configurator) {
// 	types.RegisterMsgServer(cfg.MsgServer(), rpsKeeper.NewMsgServerImpl(am.keeper))
// 	types.RegisterQueryServer(cfg.QueryServer(), rpsKeeper.NewQueryServerImpl(am.keeper))
// }

// ********************* IMPLEMENT HasGenesis INTERFACE ******************
// DefaultGenesis returns default genesis state as raw bytes for the module.
func (AppModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.NewGenesisState())
}

// ValidateGenesis performs genesis state validation for the circuit module.
func (AppModule) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	err := cdc.UnmarshalJSON(bz, &data) // Unmarshal state data
	if err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}

	return data.Validate() // Validate state data
}

// InitGenesis performs genesis initialization for the checkers module.
// It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	if err := am.keeper.InitGenesis(ctx, &genesisState); err != nil {
		panic(fmt.Sprintf("failed to initialize %s genesis state: %v", types.ModuleName, err))
	}
}

// ExportGenesis returns the exported genesis state as raw bytes for the circuit
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs, err := am.keeper.ExportGenesis(ctx)
	if err != nil {
		panic(fmt.Sprintf("failed to export %s genesis state: %v", types.ModuleName, err))
	}
	return cdc.MustMarshalJSON(gs)
}

// ***********************************************************************

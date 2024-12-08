package rpsKeeper

import (
	"fmt"

	"challenge/x/rps/types"
	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
)

type Keeper struct {
	cdc          codec.BinaryCodec // Serialize and deserialize
	addressCodec address.Codec

	// authority is the address capable of executing a MsgUpdateParams and other authority-gated message.
	// typically, this should be the x/gov module account.
	authority string

	// state management
	Schema   collections.Schema
	Students collections.Map[string, types.Student] // Map [Id] => Student
}

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService, authority string) Keeper {
	// Decode authority and check
	_, err := addressCodec.StringToBytes(authority)
	if err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}

	sb := collections.NewSchemaBuilder(storeService) // instance used to define the keeper
	k := Keeper{
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,
	}

	schema, err := sb.Build() // Build the whole squema
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k *Keeper) GetAuthority() string {
	return k.authority
}

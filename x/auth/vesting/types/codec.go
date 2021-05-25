package types

import (
	"github.com/soominhyunwoo/chain-sdk/codec"
	"github.com/soominhyunwoo/chain-sdk/codec/types"
	sdk "github.com/soominhyunwoo/chain-sdk/types"
	"github.com/soominhyunwoo/chain-sdk/types/msgservice"
	authtypes "github.com/soominhyunwoo/chain-sdk/x/auth/types"
	"github.com/soominhyunwoo/chain-sdk/x/auth/vesting/exported"
)

// RegisterLegacyAminoCodec registers the vesting interfaces and concrete types on the
// provided LegacyAmino codec. These types are used for Amino JSON serialization
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*exported.VestingAccount)(nil), nil)
	cdc.RegisterConcrete(&BaseVestingAccount{}, "soominhyunwoo-sdk/BaseVestingAccount", nil)
	cdc.RegisterConcrete(&ContinuousVestingAccount{}, "soominhyunwoo-sdk/ContinuousVestingAccount", nil)
	cdc.RegisterConcrete(&DelayedVestingAccount{}, "soominhyunwoo-sdk/DelayedVestingAccount", nil)
	cdc.RegisterConcrete(&PeriodicVestingAccount{}, "soominhyunwoo-sdk/PeriodicVestingAccount", nil)
	cdc.RegisterConcrete(&PermanentLockedAccount{}, "soominhyunwoo-sdk/PermanentLockedAccount", nil)
}

// RegisterInterface associates protoName with AccountI and VestingAccount
// Interfaces and creates a registry of it's concrete implementations
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"soominhyunwoo.vesting.v1beta1.VestingAccount",
		(*exported.VestingAccount)(nil),
		&ContinuousVestingAccount{},
		&DelayedVestingAccount{},
		&PeriodicVestingAccount{},
		&PermanentLockedAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&BaseVestingAccount{},
		&DelayedVestingAccount{},
		&ContinuousVestingAccount{},
		&PeriodicVestingAccount{},
		&PermanentLockedAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&BaseVestingAccount{},
		&DelayedVestingAccount{},
		&ContinuousVestingAccount{},
		&PeriodicVestingAccount{},
		&PermanentLockedAccount{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateVestingAccount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var amino = codec.NewLegacyAmino()

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}

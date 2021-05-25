package types

import (
	"github.com/soominhyunwoo/chain-sdk/codec"
	"github.com/soominhyunwoo/chain-sdk/codec/types"
	cryptocodec "github.com/soominhyunwoo/chain-sdk/crypto/codec"
	sdk "github.com/soominhyunwoo/chain-sdk/types"
	"github.com/soominhyunwoo/chain-sdk/types/msgservice"
	govtypes "github.com/soominhyunwoo/chain-sdk/x/gov/types"
)

// RegisterLegacyAminoCodec registers the necessary x/distribution interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgWithdrawDelegatorReward{}, "soominhyunwoo-sdk/MsgWithdrawDelegationReward", nil)
	cdc.RegisterConcrete(&MsgWithdrawValidatorCommission{}, "soominhyunwoo-sdk/MsgWithdrawValidatorCommission", nil)
	cdc.RegisterConcrete(&MsgSetWithdrawAddress{}, "soominhyunwoo-sdk/MsgModifyWithdrawAddress", nil)
	cdc.RegisterConcrete(&MsgFundCommunityPool{}, "soominhyunwoo-sdk/MsgFundCommunityPool", nil)
	cdc.RegisterConcrete(&CommunityPoolSpendProposal{}, "soominhyunwoo-sdk/CommunityPoolSpendProposal", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgWithdrawDelegatorReward{},
		&MsgWithdrawValidatorCommission{},
		&MsgSetWithdrawAddress{},
		&MsgFundCommunityPool{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&CommunityPoolSpendProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/distribution module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding as Amino
	// is still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/distribution and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

package types

import (
	"github.com/soominhyunwoo/chain-sdk/codec"
	"github.com/soominhyunwoo/chain-sdk/codec/types"
)

const (
	// MsgInterfaceProtoName defines the protobuf name of the soominhyunwoo Msg interface
	MsgInterfaceProtoName = "soominhyunwoo.base.v1beta1.Msg"
)

// RegisterLegacyAminoCodec registers the sdk message type.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*Msg)(nil), nil)
	cdc.RegisterInterface((*Tx)(nil), nil)
}

// RegisterInterfaces registers the sdk message type.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(MsgInterfaceProtoName, (*Msg)(nil))
}

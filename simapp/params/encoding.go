package params

import (
	"github.com/soominhyunwoo/chain-sdk/client"
	"github.com/soominhyunwoo/chain-sdk/codec"
	"github.com/soominhyunwoo/chain-sdk/codec/types"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

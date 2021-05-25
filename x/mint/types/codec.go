package types

import (
	"github.com/soominhyunwoo/chain-sdk/codec"
	cryptocodec "github.com/soominhyunwoo/chain-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

package rosetta

import (
	"github.com/soominhyunwoo/chain-sdk/codec"
	codectypes "github.com/soominhyunwoo/chain-sdk/codec/types"
	cryptocodec "github.com/soominhyunwoo/chain-sdk/crypto/codec"
	authcodec "github.com/soominhyunwoo/chain-sdk/x/auth/types"
	bankcodec "github.com/soominhyunwoo/chain-sdk/x/bank/types"
)

// MakeCodec generates the codec required to interact
// with the soominhyunwoo APIs used by the rosetta gateway
func MakeCodec() (*codec.ProtoCodec, codectypes.InterfaceRegistry) {
	ir := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(ir)

	authcodec.RegisterInterfaces(ir)
	bankcodec.RegisterInterfaces(ir)
	cryptocodec.RegisterInterfaces(ir)

	return cdc, ir
}

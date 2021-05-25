package legacytx

import (
	"github.com/soominhyunwoo/chain-sdk/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(StdTx{}, "soominhyunwoo-sdk/StdTx", nil)
}

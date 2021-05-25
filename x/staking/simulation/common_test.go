package simulation_test

import (
	"math/big"

	sdk "github.com/soominhyunwoo/chain-sdk/types"
)

func init() {
	sdk.DefaultPowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
}

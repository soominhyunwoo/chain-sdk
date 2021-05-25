package keeper

import (
	abci "github.com/soominhyunwoo/tendermint/abci/types"
	tmproto "github.com/soominhyunwoo/tendermint/proto/tendermint/types"

	"github.com/soominhyunwoo/chain-sdk/baseapp"
	"github.com/soominhyunwoo/chain-sdk/x/params/types"
)

// ConsensusParamsKeyTable returns an x/params module keyTable to be used in
// the BaseApp's ParamStore. The KeyTable registers the types along with the
// standard validation functions. Applications can choose to adopt this KeyTable
// or provider their own when the existing validation functions do not suite their
// needs.
func ConsensusParamsKeyTable() types.KeyTable {
	return types.NewKeyTable(
		types.NewParamSetPair(
			baseapp.ParamStoreKeyBlockParams, abci.BlockParams{}, baseapp.ValidateBlockParams,
		),
		types.NewParamSetPair(
			baseapp.ParamStoreKeyEvidenceParams, tmproto.EvidenceParams{}, baseapp.ValidateEvidenceParams,
		),
		types.NewParamSetPair(
			baseapp.ParamStoreKeyValidatorParams, tmproto.ValidatorParams{}, baseapp.ValidateValidatorParams,
		),
	)
}

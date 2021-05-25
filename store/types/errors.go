package types

import (
	sdkerrors "github.com/soominhyunwoo/chain-sdk/types/errors"
)

const StoreCodespace = "store"

var (
	ErrInvalidProof = sdkerrors.Register(StoreCodespace, 2, "invalid proof")
)

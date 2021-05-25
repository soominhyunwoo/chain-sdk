package client

import (
	sdk "github.com/soominhyunwoo/chain-sdk/types"
	sdkerrors "github.com/soominhyunwoo/chain-sdk/types/errors"
	govtypes "github.com/soominhyunwoo/chain-sdk/x/gov/types"
	"github.com/soominhyunwoo/chain-sdk/x/ibc/core/02-client/keeper"
	"github.com/soominhyunwoo/chain-sdk/x/ibc/core/02-client/types"
)

// NewClientUpdateProposalHandler defines the client update proposal handler
func NewClientUpdateProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.ClientUpdateProposal:
			return k.ClientUpdateProposal(ctx, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized ibc proposal content type: %T", c)
		}
	}
}

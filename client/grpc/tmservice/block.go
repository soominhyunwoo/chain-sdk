package tmservice

import (
	"context"

	ctypes "github.com/soominhyunwoo/tendermint/rpc/core/types"

	"github.com/soominhyunwoo/chain-sdk/client"
)

func getBlock(ctx context.Context, clientCtx client.Context, height *int64) (*ctypes.ResultBlock, error) {
	// get the node
	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	return node.Block(ctx, height)
}

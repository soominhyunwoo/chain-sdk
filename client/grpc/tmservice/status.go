package tmservice

import (
	"context"

	ctypes "github.com/soominhyunwoo/tendermint/rpc/core/types"

	"github.com/soominhyunwoo/chain-sdk/client"
)

func getNodeStatus(ctx context.Context, clientCtx client.Context) (*ctypes.ResultStatus, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return &ctypes.ResultStatus{}, err
	}
	return node.Status(ctx)
}

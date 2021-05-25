package testutil

import (
	"github.com/soominhyunwoo/chain-sdk/testutil"
	clitestutil "github.com/soominhyunwoo/chain-sdk/testutil/cli"
	"github.com/soominhyunwoo/chain-sdk/testutil/network"
	"github.com/soominhyunwoo/chain-sdk/x/authz/client/cli"
)

func ExecGrant(val *network.Validator, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	clientCtx := val.ClientCtx
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}

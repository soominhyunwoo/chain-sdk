package testutil

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	tmcfg "github.com/soominhyunwoo/tendermint/config"
	"github.com/soominhyunwoo/tendermint/libs/cli"
	"github.com/soominhyunwoo/tendermint/libs/log"

	"github.com/soominhyunwoo/chain-sdk/client"
	"github.com/soominhyunwoo/chain-sdk/codec"
	"github.com/soominhyunwoo/chain-sdk/server"
	"github.com/soominhyunwoo/chain-sdk/testutil"
	"github.com/soominhyunwoo/chain-sdk/types/module"
	genutilcli "github.com/soominhyunwoo/chain-sdk/x/genutil/client/cli"
)

func ExecInitCmd(testMbm module.BasicManager, home string, cdc codec.JSONCodec) error {
	logger := log.NewNopLogger()
	cfg, err := CreateDefaultTendermintConfig(home)
	if err != nil {
		return err
	}

	cmd := genutilcli.InitCmd(testMbm, home)
	serverCtx := server.NewContext(viper.New(), cfg, logger)
	clientCtx := client.Context{}.WithJSONCodec(cdc).WithHomeDir(home)

	_, out := testutil.ApplyMockIO(cmd)
	clientCtx = clientCtx.WithOutput(out)

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
	ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)

	cmd.SetArgs([]string{"appnode-test", fmt.Sprintf("--%s=%s", cli.HomeFlag, home)})

	return cmd.ExecuteContext(ctx)
}

func CreateDefaultTendermintConfig(rootDir string) (*tmcfg.Config, error) {
	conf := tmcfg.DefaultConfig()
	conf.SetRoot(rootDir)
	tmcfg.EnsureRoot(rootDir)

	if err := conf.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("error in config file: %v", err)
	}

	return conf, nil
}

package cmd_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/soominhyunwoo/tendermint/libs/log"

	"github.com/soominhyunwoo/chain-sdk/client"
	"github.com/soominhyunwoo/chain-sdk/client/flags"
	"github.com/soominhyunwoo/chain-sdk/server"
	"github.com/soominhyunwoo/chain-sdk/simapp"
	simcmd "github.com/soominhyunwoo/chain-sdk/simapp/simd/cmd"
	"github.com/soominhyunwoo/chain-sdk/testutil/testdata"
	"github.com/soominhyunwoo/chain-sdk/types/module"
	"github.com/soominhyunwoo/chain-sdk/x/genutil"
	genutiltest "github.com/soominhyunwoo/chain-sdk/x/genutil/client/testutil"
)

var testMbm = module.NewBasicManager(genutil.AppModuleBasic{})

func TestAddGenesisAccountCmd(t *testing.T) {
	_, _, addr1 := testdata.KeyTestPubAddr()
	tests := []struct {
		name      string
		addr      string
		denom     string
		expectErr bool
	}{
		{
			name:      "invalid address",
			addr:      "",
			denom:     "1000atom",
			expectErr: true,
		},
		{
			name:      "valid address",
			addr:      addr1.String(),
			denom:     "1000atom",
			expectErr: false,
		},
		{
			name:      "multiple denoms",
			addr:      addr1.String(),
			denom:     "1000atom, 2000stake",
			expectErr: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			home := t.TempDir()
			logger := log.NewNopLogger()
			cfg, err := genutiltest.CreateDefaultTendermintConfig(home)
			require.NoError(t, err)

			appCodec := simapp.MakeTestEncodingConfig().Marshaler
			err = genutiltest.ExecInitCmd(testMbm, home, appCodec)
			require.NoError(t, err)

			serverCtx := server.NewContext(viper.New(), cfg, logger)
			clientCtx := client.Context{}.WithJSONCodec(appCodec).WithHomeDir(home)

			ctx := context.Background()
			ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
			ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)

			cmd := simcmd.AddGenesisAccountCmd(home)
			cmd.SetArgs([]string{
				tc.addr,
				tc.denom,
				fmt.Sprintf("--%s=home", flags.FlagHome)})

			if tc.expectErr {
				require.Error(t, cmd.ExecuteContext(ctx))
			} else {
				require.NoError(t, cmd.ExecuteContext(ctx))
			}
		})
	}
}

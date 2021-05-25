package keeper_test

import (
	tmproto "github.com/soominhyunwoo/tendermint/proto/tendermint/types"

	"github.com/soominhyunwoo/chain-sdk/simapp"
	sdk "github.com/soominhyunwoo/chain-sdk/types"
	authtypes "github.com/soominhyunwoo/chain-sdk/x/auth/types"
)

// returns context and app with params set on account keeper
func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())

	return app, ctx
}

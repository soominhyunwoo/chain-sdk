package keeper_test

import (
	tmproto "github.com/soominhyunwoo/tendermint/proto/tendermint/types"

	"github.com/soominhyunwoo/chain-sdk/simapp"
	sdk "github.com/soominhyunwoo/chain-sdk/types"
	"github.com/soominhyunwoo/chain-sdk/x/mint/types"
)

// returns context and an app with updated mint keeper
func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.MintKeeper.SetParams(ctx, types.DefaultParams())
	app.MintKeeper.SetMinter(ctx, types.DefaultInitialMinter())

	return app, ctx
}

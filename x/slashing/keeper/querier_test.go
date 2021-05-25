package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	abci "github.com/soominhyunwoo/tendermint/abci/types"
	tmproto "github.com/soominhyunwoo/tendermint/proto/tendermint/types"

	"github.com/soominhyunwoo/chain-sdk/codec"
	"github.com/soominhyunwoo/chain-sdk/simapp"
	"github.com/soominhyunwoo/chain-sdk/x/slashing/keeper"
	"github.com/soominhyunwoo/chain-sdk/x/slashing/testslashing"
	"github.com/soominhyunwoo/chain-sdk/x/slashing/types"
)

func TestNewQuerier(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	app.SlashingKeeper.SetParams(ctx, testslashing.TestParams())
	legacyQuerierCdc := codec.NewAminoCodec(app.LegacyAmino())
	querier := keeper.NewQuerier(app.SlashingKeeper, legacyQuerierCdc.LegacyAmino)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	_, err := querier(ctx, []string{types.QueryParameters}, query)
	require.NoError(t, err)
}

func TestQueryParams(t *testing.T) {
	cdc := codec.NewLegacyAmino()
	legacyQuerierCdc := codec.NewAminoCodec(cdc)
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	app.SlashingKeeper.SetParams(ctx, testslashing.TestParams())

	querier := keeper.NewQuerier(app.SlashingKeeper, legacyQuerierCdc.LegacyAmino)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	var params types.Params

	res, err := querier(ctx, []string{types.QueryParameters}, query)
	require.NoError(t, err)

	err = cdc.UnmarshalJSON(res, &params)
	require.NoError(t, err)
	require.Equal(t, app.SlashingKeeper.GetParams(ctx), params)
}

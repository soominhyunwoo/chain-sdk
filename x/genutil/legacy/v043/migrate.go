package v043

import (
	"github.com/soominhyunwoo/chain-sdk/client"
	"github.com/soominhyunwoo/chain-sdk/x/genutil/types"
	v040gov "github.com/soominhyunwoo/chain-sdk/x/gov/legacy/v040"
	v043gov "github.com/soominhyunwoo/chain-sdk/x/gov/legacy/v043"
)

// Migrate migrates exported state from v0.40 to a v0.43 genesis state.
func Migrate(appState types.AppMap, clientCtx client.Context) types.AppMap {
	// Migrate x/gov.
	if appState[v040gov.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var oldGovState v040gov.GenesisState
		clientCtx.JSONCodec.MustUnmarshalJSON(appState[v040gov.ModuleName], &oldGovState)

		// delete deprecated x/gov genesis state
		delete(appState, v040gov.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v043gov.ModuleName] = clientCtx.JSONCodec.MustMarshalJSON(v043gov.MigrateJSON(&oldGovState))
	}

	return appState
}

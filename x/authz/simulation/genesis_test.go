package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/soominhyunwoo/chain-sdk/simapp"
	"github.com/soominhyunwoo/chain-sdk/types/module"
	simtypes "github.com/soominhyunwoo/chain-sdk/types/simulation"
	"github.com/soominhyunwoo/chain-sdk/x/authz"
	"github.com/soominhyunwoo/chain-sdk/x/authz/simulation"
)

func TestRandomizedGenState(t *testing.T) {
	app := simapp.Setup(false)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          app.AppCodec(),
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: 1000,
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)
	var authzGenesis authz.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[authz.ModuleName], &authzGenesis)

	require.Len(t, authzGenesis.Authorization, len(simState.Accounts)-1)
}

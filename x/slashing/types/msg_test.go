package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/soominhyunwoo/chain-sdk/types"
)

func TestMsgUnjailGetSignBytes(t *testing.T) {
	addr := sdk.AccAddress("abcd")
	msg := NewMsgUnjail(sdk.ValAddress(addr))
	bytes := msg.GetSignBytes()
	require.Equal(
		t,
		`{"type":"soominhyunwoo-sdk/MsgUnjail","value":{"address":"soominhyunwoovaloper1v93xxeqhg9nn6"}}`,
		string(bytes),
	)
}

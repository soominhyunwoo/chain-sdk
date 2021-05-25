package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/soominhyunwoo/chain-sdk/codec"
	cryptoAmino "github.com/soominhyunwoo/chain-sdk/crypto/codec"
	"github.com/soominhyunwoo/chain-sdk/testutil/testdata"
	sdk "github.com/soominhyunwoo/chain-sdk/types"
	"github.com/soominhyunwoo/chain-sdk/x/auth/legacy/legacytx"
	"github.com/soominhyunwoo/chain-sdk/x/auth/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "soominhyunwoo-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, testutil.NewTxConfigTestSuite(txGen))
}

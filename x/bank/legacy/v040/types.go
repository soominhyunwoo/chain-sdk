package v040

import (
	"github.com/golang/protobuf/proto"

	codectypes "github.com/soominhyunwoo/chain-sdk/codec/types"
	"github.com/soominhyunwoo/chain-sdk/x/bank/types"
)

// SupplyI defines an inflationary supply interface for modules that handle
// token supply.
// It is copy-pasted from:
// https://github.com/soominhyunwoo/chain-sdk/blob/v042.3/x/bank/exported/exported.go
// where we stripped off the unnecessary methods.
//
// It is used in the migration script, because we save this interface as an Any
// in the supply state.
//
// Deprecated.
type SupplyI interface {
	proto.Message
}

// RegisterInterfaces registers interfaces required for the v0.40 migrations.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"soominhyunwoo.bank.v1beta1.SupplyI",
		(*SupplyI)(nil),
		&types.Supply{},
	)
}

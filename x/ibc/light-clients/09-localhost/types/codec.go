package types

import (
	codectypes "github.com/soominhyunwoo/chain-sdk/codec/types"
	"github.com/soominhyunwoo/chain-sdk/x/ibc/core/exported"
)

// RegisterInterfaces register the ibc interfaces submodule implementations to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
}

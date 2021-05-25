package authz

import (
	types "github.com/soominhyunwoo/chain-sdk/codec/types"
	sdk "github.com/soominhyunwoo/chain-sdk/types"
	"github.com/soominhyunwoo/chain-sdk/types/msgservice"
)

// RegisterInterfaces registers the interfaces types with the interface registry
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGrant{},
		&MsgRevoke{},
		&MsgExec{},
	)

	registry.RegisterInterface(
		"soominhyunwoo.v1beta1.Authorization",
		(*Authorization)(nil),
		&GenericAuthorization{},
	)

	msgservice.RegisterMsgServiceDesc(registry, MsgServiceDesc())
}

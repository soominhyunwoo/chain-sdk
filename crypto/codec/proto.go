package codec

import (
	codectypes "github.com/soominhyunwoo/chain-sdk/codec/types"
	"github.com/soominhyunwoo/chain-sdk/crypto/keys/ed25519"
	"github.com/soominhyunwoo/chain-sdk/crypto/keys/multisig"
	"github.com/soominhyunwoo/chain-sdk/crypto/keys/secp256k1"
	"github.com/soominhyunwoo/chain-sdk/crypto/keys/secp256r1"
	cryptotypes "github.com/soominhyunwoo/chain-sdk/crypto/types"
)

// RegisterInterfaces registers the sdk.Tx interface.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	var pk *cryptotypes.PubKey
	registry.RegisterInterface("soominhyunwoo.crypto.PubKey", pk)
	registry.RegisterImplementations(pk, &ed25519.PubKey{})
	registry.RegisterImplementations(pk, &secp256k1.PubKey{})
	registry.RegisterImplementations(pk, &multisig.LegacyAminoPubKey{})
	secp256r1.RegisterInterfaces(registry)
}

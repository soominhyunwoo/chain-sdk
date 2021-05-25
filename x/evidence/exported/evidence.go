package exported

import (
	"github.com/gogo/protobuf/proto"
	tmbytes "github.com/soominhyunwoo/tendermint/libs/bytes"

	sdk "github.com/soominhyunwoo/chain-sdk/types"
)

// Evidence defines the contract which concrete evidence types of misbehavior
// must implement.
type Evidence interface {
	proto.Message

	Route() string
	Type() string
	String() string
	Hash() tmbytes.HexBytes
	ValidateBasic() error

	// Height at which the infraction occurred
	GetHeight() int64
}

// ValidatorEvidence extends Evidence interface to define contract
// for evidence against malicious validators
type ValidatorEvidence interface {
	Evidence

	// The consensus address of the malicious validator at time of infraction
	GetConsensusAddress() sdk.ConsAddress

	// The total power of the malicious validator at time of infraction
	GetValidatorPower() int64

	// The total validator set power at time of infraction
	GetTotalPower() int64
}

// MsgSubmitEvidenceI defines the specific interface a concrete message must
// implement in order to process submitted evidence. The concrete MsgSubmitEvidence
// must be defined at the application-level.
type MsgSubmitEvidenceI interface {
	sdk.Msg

	GetEvidence() Evidence
	GetSubmitter() sdk.AccAddress
}

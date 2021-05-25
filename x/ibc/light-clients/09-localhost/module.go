package localhost

import (
	"github.com/soominhyunwoo/chain-sdk/x/ibc/light-clients/09-localhost/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}

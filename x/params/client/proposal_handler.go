package client

import (
	govclient "github.com/soominhyunwoo/chain-sdk/x/gov/client"
	"github.com/soominhyunwoo/chain-sdk/x/params/client/cli"
	"github.com/soominhyunwoo/chain-sdk/x/params/client/rest"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd, rest.ProposalRESTHandler)

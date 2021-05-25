package client

import (
	"github.com/soominhyunwoo/chain-sdk/x/distribution/client/cli"
	"github.com/soominhyunwoo/chain-sdk/x/distribution/client/rest"
	govclient "github.com/soominhyunwoo/chain-sdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)

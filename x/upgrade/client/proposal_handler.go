package client

import (
	govclient "github.com/soominhyunwoo/chain-sdk/x/gov/client"
	"github.com/soominhyunwoo/chain-sdk/x/upgrade/client/cli"
	"github.com/soominhyunwoo/chain-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal, rest.ProposalCancelRESTHandler)

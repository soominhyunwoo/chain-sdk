// +build norace

package testutil

import (
	"testing"

	"github.com/soominhyunwoo/chain-sdk/testutil/network"

	"github.com/stretchr/testify/suite"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 3
	suite.Run(t, NewIntegrationTestSuite(cfg))
}
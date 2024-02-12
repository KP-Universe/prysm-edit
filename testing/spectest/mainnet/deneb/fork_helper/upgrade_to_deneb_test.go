package fork_helper

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/deneb/fork"
)

func TestMainnet_UpgradeToDeneb(t *testing.T) {
	fork.RunUpgradeToDeneb(t, "mainnet")
}

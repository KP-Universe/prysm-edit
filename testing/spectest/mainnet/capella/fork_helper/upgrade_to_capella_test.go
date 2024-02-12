package fork_helper

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/capella/fork"
)

func TestMainnet_Capella_UpgradeToCapella(t *testing.T) {
	fork.RunUpgradeToCapella(t, "mainnet")
}

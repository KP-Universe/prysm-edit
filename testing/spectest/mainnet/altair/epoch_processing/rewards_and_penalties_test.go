package epoch_processing

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/altair/epoch_processing"
)

func TestMainnet_Altair_EpochProcessing_RewardsAndPenalties(t *testing.T) {
	epoch_processing.RunRewardsAndPenaltiesTests(t, "mainnet")
}

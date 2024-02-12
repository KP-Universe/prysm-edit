package finality

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/phase0/finality"
)

func TestMainnet_Phase0_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "mainnet")
}

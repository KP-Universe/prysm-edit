package random

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/bellatrix/sanity"
)

func TestMainnet_Bellatrix_Random(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "mainnet", "random/random/pyspec_tests")
}

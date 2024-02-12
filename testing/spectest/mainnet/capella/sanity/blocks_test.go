package sanity

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/capella/sanity"
)

func TestMainnet_Capella_Sanity_Blocks(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "mainnet", "sanity/blocks/pyspec_tests")
}

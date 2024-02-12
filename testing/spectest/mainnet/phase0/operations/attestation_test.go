package operations

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/phase0/operations"
)

func TestMainnet_Phase0_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "mainnet")
}

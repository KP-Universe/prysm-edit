package operations

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/altair/operations"
)

func TestMainnet_Altair_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "mainnet")
}

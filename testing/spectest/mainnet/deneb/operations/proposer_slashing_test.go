package operations

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/deneb/operations"
)

func TestMainnet_Deneb_Operations_ProposerSlashing(t *testing.T) {
	operations.RunProposerSlashingTest(t, "mainnet")
}

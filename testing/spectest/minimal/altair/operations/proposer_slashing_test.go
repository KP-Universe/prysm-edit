package operations

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/altair/operations"
)

func TestMinimal_Altair_Operations_ProposerSlashing(t *testing.T) {
	operations.RunProposerSlashingTest(t, "minimal")
}

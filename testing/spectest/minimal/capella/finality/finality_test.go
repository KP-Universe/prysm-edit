package finality

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/capella/finality"
)

func TestMinimal_Capella_Finality(t *testing.T) {
	finality.RunFinalityTest(t, "minimal")
}

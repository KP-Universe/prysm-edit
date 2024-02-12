package random

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/capella/sanity"
)

func TestMinimal_Capella_Random(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "minimal", "random/random/pyspec_tests")
}

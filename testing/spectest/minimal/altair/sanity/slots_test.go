package sanity

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/altair/sanity"
)

func TestMinimal_Altair_Sanity_Slots(t *testing.T) {
	sanity.RunSlotProcessingTests(t, "minimal")
}

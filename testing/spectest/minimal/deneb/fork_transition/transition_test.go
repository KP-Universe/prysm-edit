package fork_transition

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/deneb/fork"
)

func TestMinimal_Deneb_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "minimal")
}

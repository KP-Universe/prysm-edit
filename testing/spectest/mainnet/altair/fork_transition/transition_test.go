package fork_transition

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/altair/fork"
)

func TestMainnet_Altair_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "mainnet")
}

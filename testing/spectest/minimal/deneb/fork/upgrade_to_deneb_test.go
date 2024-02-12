package fork

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/deneb/fork"
)

func TestMinimal_UpgradeToDeneb(t *testing.T) {
	fork.RunUpgradeToDeneb(t, "minimal")
}

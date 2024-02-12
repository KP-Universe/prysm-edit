package forkchoice

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/runtime/version"
	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/common/forkchoice"
)

func TestMainnet_Capella_Forkchoice(t *testing.T) {
	forkchoice.Run(t, "mainnet", version.Capella)
}

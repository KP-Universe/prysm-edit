package operations

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/bellatrix/operations"
)

func TestMainnet_Bellatrix_Operations_VoluntaryExit(t *testing.T) {
	operations.RunVoluntaryExitTest(t, "mainnet")
}

package operations

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/bellatrix/operations"
)

func TestMinimal_Bellatrix_Operations_PayloadExecution(t *testing.T) {
	operations.RunExecutionPayloadTest(t, "minimal")
}

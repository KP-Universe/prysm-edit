package operations

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/capella/operations"
)

func TestMinimal_Capella_Operations_PayloadExecution(t *testing.T) {
	operations.RunExecutionPayloadTest(t, "minimal")
}

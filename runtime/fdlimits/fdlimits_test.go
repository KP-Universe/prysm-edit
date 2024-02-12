package fdlimits_test

import (
	"testing"

	gethLimit "github.com/KP-Universe/go-kpu/common/fdlimit"
	"github.com/KP-Universe/prysm/v4/runtime/fdlimits"
	"github.com/KP-Universe/prysm/v4/testing/assert"
)

func TestSetMaxFdLimits(t *testing.T) {
	assert.NoError(t, fdlimits.SetMaxFdLimits())

	curr, err := gethLimit.Current()
	assert.NoError(t, err)

	max, err := gethLimit.Maximum()
	assert.NoError(t, err)

	assert.Equal(t, max, curr, "current and maximum file descriptor limits do not match up.")

}

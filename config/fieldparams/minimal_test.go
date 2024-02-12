//go:build minimal

package field_params_test

import (
	"testing"

	fieldparams "github.com/KP-Universe/prysm/v4/config/fieldparams"
	"github.com/KP-Universe/prysm/v4/config/params"
	"github.com/KP-Universe/prysm/v4/testing/require"
)

func TestFieldParametersValues(t *testing.T) {
	params.SetupTestConfigCleanup(t)
	min := params.MinimalSpecConfig().Copy()
	params.OverrideBeaconConfig(min)
	require.Equal(t, "minimal", fieldparams.Preset)
	testFieldParametersMatchConfig(t)
}

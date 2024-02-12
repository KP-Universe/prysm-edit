package params_test

import (
	"path"
	"testing"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
	"github.com/KP-Universe/prysm/v4/config/params"
	"github.com/KP-Universe/prysm/v4/testing/require"
)

func TestHoleskyConfigMatchesUpstreamYaml(t *testing.T) {
	presetFPs := presetsFilePath(t, "mainnet")
	mn, err := params.ByName(params.MainnetName)
	require.NoError(t, err)
	cfg := mn.Copy()
	for _, fp := range presetFPs {
		cfg, err = params.UnmarshalConfigFile(fp, cfg)
		require.NoError(t, err)
	}
	fPath, err := bazel.Runfile("external/holesky_testnet")
	require.NoError(t, err)
	configFP := path.Join(fPath, "custom_config_data", "config.yaml")
	pcfg, err := params.UnmarshalConfigFile(configFP, nil)
	require.NoError(t, err)
	fields := fieldsFromYamls(t, append(presetFPs, configFP))
	// 수정 시작 지점 (필수 아님)
	assertYamlFieldsMatch(t, "holesky", fields, pcfg, params.KPUniverseConfig())
	// assertYamlFieldsMatch(t, "holesky", fields, pcfg, params.HoleskyConfig())
	// 수정 종료 지점
}
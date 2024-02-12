package operations

import (
	"context"
	"path"
	"testing"

	"github.com/golang/snappy"
	"github.com/KP-Universe/prysm/v4/beacon-chain/core/blocks"
	"github.com/KP-Universe/prysm/v4/beacon-chain/core/validators"
	"github.com/KP-Universe/prysm/v4/beacon-chain/state"
	"github.com/KP-Universe/prysm/v4/consensus-types/interfaces"
	ethpb "github.com/KP-Universe/prysm/v4/proto/prysm/v1alpha1"
	"github.com/KP-Universe/prysm/v4/testing/require"
	"github.com/KP-Universe/prysm/v4/testing/spectest/utils"
	"github.com/KP-Universe/prysm/v4/testing/util"
)

func RunProposerSlashingTest(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))
	testFolders, testsFolderPath := utils.TestFolders(t, config, "capella", "operations/proposer_slashing/pyspec_tests")
	if len(testFolders) == 0 {
		t.Fatalf("No test folders found for %s/%s/%s", config, "capella", "operations/proposer_slashing/pyspec_tests")
	}
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			proposerSlashingFile, err := util.BazelFileBytes(folderPath, "proposer_slashing.ssz_snappy")
			require.NoError(t, err)
			proposerSlashingSSZ, err := snappy.Decode(nil /* dst */, proposerSlashingFile)
			require.NoError(t, err, "Failed to decompress")
			proposerSlashing := &ethpb.ProposerSlashing{}
			require.NoError(t, proposerSlashing.UnmarshalSSZ(proposerSlashingSSZ), "Failed to unmarshal")

			body := &ethpb.BeaconBlockBodyCapella{ProposerSlashings: []*ethpb.ProposerSlashing{proposerSlashing}}
			RunBlockOperationTest(t, folderPath, body, func(ctx context.Context, s state.BeaconState, b interfaces.ReadOnlySignedBeaconBlock) (state.BeaconState, error) {
				return blocks.ProcessProposerSlashings(ctx, s, b.Block().Body().ProposerSlashings(), validators.SlashValidator)
			})
		})
	}
}

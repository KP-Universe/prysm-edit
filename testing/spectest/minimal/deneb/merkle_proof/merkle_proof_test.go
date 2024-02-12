package merkle_proof

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/spectest/shared/deneb/merkle_proof"
)

func TestMainnet_Deneb_MerkleProof(t *testing.T) {
	merkle_proof.RunMerkleProofTests(t, "minimal")
}

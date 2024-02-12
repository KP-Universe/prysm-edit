package attestations

import (
	"github.com/KP-Universe/prysm/v4/beacon-chain/operations/attestations/kv"
)

var _ Pool = (*kv.AttCaches)(nil)

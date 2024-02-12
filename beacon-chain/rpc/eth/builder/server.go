package builder

import (
	"github.com/KP-Universe/prysm/v4/beacon-chain/blockchain"
	"github.com/KP-Universe/prysm/v4/beacon-chain/rpc/lookup"
)

type Server struct {
	FinalizationFetcher   blockchain.FinalizationFetcher
	OptimisticModeFetcher blockchain.OptimisticModeFetcher
	Stater                lookup.Stater
}

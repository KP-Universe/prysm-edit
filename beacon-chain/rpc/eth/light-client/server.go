package lightclient

import (
	"github.com/KP-Universe/prysm/v4/beacon-chain/blockchain"
	"github.com/KP-Universe/prysm/v4/beacon-chain/rpc/lookup"
)

type Server struct {
	Blocker     lookup.Blocker
	Stater      lookup.Stater
	HeadFetcher blockchain.HeadFetcher
}

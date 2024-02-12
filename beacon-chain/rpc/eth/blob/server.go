package blob

import (
	"github.com/KP-Universe/prysm/v4/beacon-chain/rpc/lookup"
)

type Server struct {
	Blocker lookup.Blocker
}

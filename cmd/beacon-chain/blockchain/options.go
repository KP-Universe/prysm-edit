package blockchaincmd

import (
	"github.com/KP-Universe/prysm/v4/beacon-chain/blockchain"
	"github.com/KP-Universe/prysm/v4/beacon-chain/core/helpers"
	"github.com/KP-Universe/prysm/v4/cmd"
	"github.com/KP-Universe/prysm/v4/cmd/beacon-chain/flags"
	"github.com/urfave/cli/v2"
)

// FlagOptions for blockchain service flag configurations.
func FlagOptions(c *cli.Context) ([]blockchain.Option, error) {
	wsp := c.String(flags.WeakSubjectivityCheckpoint.Name)
	wsCheckpt, err := helpers.ParseWeakSubjectivityInputString(wsp)
	if err != nil {
		return nil, err
	}
	maxRoutines := c.Int(cmd.MaxGoroutines.Name)
	opts := []blockchain.Option{
		blockchain.WithMaxGoroutines(maxRoutines),
		blockchain.WithWeakSubjectivityCheckpoint(wsCheckpt),
	}
	return opts, nil
}

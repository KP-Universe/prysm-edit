package deprecated

import (
	"github.com/KP-Universe/prysm/v4/cmd/prysmctl/deprecated/checkpoint"
	"github.com/urfave/cli/v2"
)

var Commands []*cli.Command

func init() {
	Commands = append(Commands, checkpoint.Commands...)
}

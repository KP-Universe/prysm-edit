package shuffle

import "github.com/KP-Universe/prysm/v4/consensus-types/primitives"

// TestCase --
type TestCase struct {
	Seed    string                      `yaml:"seed"`
	Count   uint64                      `yaml:"count"`
	Mapping []primitives.ValidatorIndex `yaml:"mapping"`
}

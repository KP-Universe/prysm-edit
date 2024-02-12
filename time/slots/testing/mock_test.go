package testing

import (
	"github.com/KP-Universe/prysm/v4/time/slots"
)

var _ slots.Ticker = (*MockTicker)(nil)

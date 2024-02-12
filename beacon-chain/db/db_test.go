package db

import "github.com/KP-Universe/prysm/v4/beacon-chain/db/kv"

var _ Database = (*kv.Store)(nil)

package endtoend

// This file contains the dependencies required for github.com/ethereum/go-ethereum/cmd/geth.
// Having these dependencies listed here helps go mod understand that these dependencies are
// necessary for end to end tests since we build go-ethereum binary for this test.
import (
	_ "github.com/KP-Universe/go-kpu/accounts"          // Required for go-ethereum e2e.
	_ "github.com/KP-Universe/go-kpu/accounts/keystore" // Required for go-ethereum e2e.
	_ "github.com/KP-Universe/go-kpu/cmd/utils"         // Required for go-ethereum e2e.
	_ "github.com/KP-Universe/go-kpu/common"            // Required for go-ethereum e2e.
	_ "github.com/KP-Universe/go-kpu/console"           // Required for go-ethereum e2e.
	_ "github.com/KP-Universe/go-kpu/eth"               // Required for go-ethereum e2e.
	_ "github.com/KP-Universe/go-kpu/eth/downloader"    // Required for go-ethereum e2e.
	_ "github.com/KP-Universe/go-kpu/ethclient"         // Required for go-ethereum e2e.
	_ "github.com/KP-Universe/go-kpu/les"               // Required for go-ethereum e2e.
	_ "github.com/KP-Universe/go-kpu/log"               // Required for go-ethereum e2e.
	_ "github.com/KP-Universe/go-kpu/metrics"           // Required for go-ethereum e2e.
	_ "github.com/KP-Universe/go-kpu/node"              // Required for go-ethereum e2e.
)

module github.com/e-money/swaptests

go 1.15

require (
	github.com/ava-labs/avalanchego v1.2.0
	github.com/ava-labs/coreth v0.3.23
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/e-money/em-ledger v0.9.4
	github.com/ethereum/go-ethereum v1.9.21
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
)

replace github.com/e-money/em-ledger => ../em-ledger

package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type SwapRequest struct {
	Id                  common.Hash
	RandomNumberHash    common.Hash
	ExpireTimestamp     int64
	SenderAddress       string
	RecipientAddress    string
	RecipientOtherChain string
	OutAmount           *big.Int
}

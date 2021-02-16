package main

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var Fixed8Decimals = big.NewInt(int64(math.Pow10(8)))

type BlockAndTxLogs struct {
	Height          int64
	BlockHash       string
	ParentBlockHash string
	BlockTime       int64
}

type EthStatus struct {
	Allowance    string `json:"allowance"`
	Erc20Balance string `json:"erc20_balance"`
	EthBalance   string `json:"eth_balance"`
}

type SwapStatus struct {
	Id               int64  `json:"id"`
	SenderAddr       string `json:"sender_addr"`
	ReceiverAddr     string `json:"receiver_addr"`
	OtherChainAddr   string `json:"other_chain_addr"`
	InAmount         string `json:"in_amount"`
	OutAmount        string `json:"out_amount"`
	RandomNumberHash string `json:"random_number_hash"`
	ExpireTimestamp  int64  `json:"expire_timestamp"`
	Height           int64  `json:"height"`
	Timestamp        int64  `json:"timestamp"`
	RandomNumber     string `json:"random_number"`
}

type FailedSwaps struct {
	TotalCount int           `json:"total_count"`
	CurPage    int           `json:"cur_page"`
	NumPerPage int           `json:"num_per_page"`
	Swaps      []*SwapStatus `json:"swaps"`
}

type ReconciliationStatus struct {
	Bep2TokenBalance          *big.Float
	Bep2TokenOutPending       *big.Float
	OtherChainTokenBalance    *big.Float
	OtherChainTokenOutPending *big.Float
}

type SwapRequest struct {
	Id                  common.Hash
	RandomNumberHash    common.Hash
	ExpireTimestamp     int64
	SenderAddress       string
	RecipientAddress    string
	RecipientOtherChain string
	OutAmount           *big.Int
}

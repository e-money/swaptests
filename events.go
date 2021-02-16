package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ava-labs/coreth/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type HTLTEvent struct {
	MsgSender        common.Address
	RecipientAddr    common.Address
	SwapId           common.Hash
	RandomNumberHash common.Hash
	Timestamp        uint64
	Bep2Addr         common.Address
	ExpireSeconds    *big.Int
	OutAmount        *big.Int
	Bep2Amount       *big.Int
}

func parseHTLTEvent(abi *abi.ABI, log *types.Log) (*HTLTEvent, error) {
	var ev HTLTEvent

	err := abi.Unpack(&ev, "HTLT", log.Data)
	if err != nil {
		return nil, err
	}

	ev.MsgSender = common.BytesToAddress(log.Topics[1].Bytes())
	ev.RecipientAddr = common.BytesToAddress(log.Topics[2].Bytes())
	ev.SwapId = common.BytesToHash(log.Topics[3].Bytes())

	fmt.Printf("sender addr: %s\n", ev.MsgSender.String())
	fmt.Printf("receiver addr: %s\n", ev.RecipientAddr.String())
	fmt.Printf("swap id: %s\n", hex.EncodeToString(ev.SwapId[:]))
	fmt.Printf("random number hash: %s\n", hex.EncodeToString(ev.RandomNumberHash[:]))
	fmt.Printf("addr: %s\n", hex.EncodeToString(ev.Bep2Addr[:]))
	fmt.Printf("timestamp: %d\n", ev.Timestamp)
	fmt.Printf("expire time: %d\n", ev.ExpireSeconds)
	fmt.Printf("erc20 amount: %d\n", ev.OutAmount)
	fmt.Printf("amount: %d\n", ev.Bep2Amount)
	return &ev, nil
}

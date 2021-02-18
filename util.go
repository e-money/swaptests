package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var randomNumber = []byte("0xaabbccddaabbccddaabbccddaabbccddaabbccddaabbccddaabbccddaabbccdd")

func assertNoError(err error, msg string) {
	if err == nil {
		return
	}

	if len(msg) == 0 {
		msg = "assertNoError failed"
	}

	panic(fmt.Sprintf("%s: %s", msg, err))
}

func assertEq(v1, v2, msg string) {
	if v1 == v2 {
		return
	}
	if len(msg) == 0 {
		msg = "assertEq failed"
	}

	panic(fmt.Sprintf("%s: %s != %s", msg, v1, v2))
}

func genRandomHash(randomNumber []byte) common.Hash {
	randomNumberPadded := make([]byte, RandomNumberLength+8)
	copy(randomNumberPadded[:RandomNumberLength], randomNumber)
	binary.BigEndian.PutUint64(randomNumberPadded[RandomNumberLength:], uint64(0))
	res := sha256.Sum256(randomNumberPadded)
	return res
}

func deriveAddressFromKey(privKey *ecdsa.PrivateKey) common.Address {
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("public key error")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address
}

// Sum returns the SHA256 of the bz.
func shaHash(bz []byte) []byte {
	h := sha256.Sum256(bz)
	return h[:]
}

func SetNetworkPrefix() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("emoney", "pub")
	config.Seal()
}

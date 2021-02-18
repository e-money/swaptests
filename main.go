package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	/******* Key must be provided for this *********************/
	/******* Producing this address */
	//ethSenderAddr = "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1"
	ethSenderAddr = "0x001d05573f994671E11d416b55cf3Cf1e5aF8dCC"

	testnetUrl = "http://localhost:9650/ext/bc/C/rpc"
	// testnetUrl = "http://localhost:8545"

	// constant addresses. no need to provide key.
	emSenderAddr = "emoney1lagqmceycrfpkyu7y6ayrk6jyvru5mkrezacpw"
	emUserAddr   = "emoney1n5ggspeff4fxc87dvmg0ematr3qzw5l4v20mdv"
	// ethRecipientAddr = "0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0"
	ethRecipientAddr = "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1"

	// testnet contracts deployed addresses
	tokenContractAddr = "0x6aeE7f22C4926f4Cca18266821B520e0dc67B294"
	swapContractAddr  = "0xCa08356d8d8f85A797eA93A893ddff874cCD7489"

	// ganache
	// tokenContractAddr = "0xCfEB869F69431e42cdB54A4F4f105C19C080A601"
	// swapContractAddr  = "0x254dffcd3277C0b1660F6d42EFbB754edaBAbC2B"

	gasLimit           = 3000000
	gasPrice           = 470000000000
	RandomNumberLength = 32
	timeSpan           = 3 * 86400 // seconds
)

// key for ethSenderAddr above. runtime arg
var ethPrvKey *ecdsa.PrivateKey

func printUsage() {
	fmt.Println("usage: ./swap -ethkey 23423d....")
}

func main() {
	var ethPrvKeyHex string
	flag.StringVar(&ethPrvKeyHex, "ethkey", "invalid default value", "eth private key")

	flag.Parse()

	var err error
	ethPrvKey, err = crypto.HexToECDSA(ethPrvKeyHex)
	if err != nil {
		printUsage()
		panic(fmt.Sprintf("invalid provided ethereum private key value:%s, error=%s", ethPrvKeyHex, err))
	}

	addr := deriveAddressFromKey(ethPrvKey)

	assertEq(ethSenderAddr, addr.String(), "runtime arg ethkey not matching ethSenderAddr")

	SetNetworkPrefix()

	if err := testSwap(); err != nil {
		os.Exit(1)
	}
}

func testSwap() error {
	e := newEthFacade(testnetUrl, common.HexToAddress(swapContractAddr), common.HexToAddress(tokenContractAddr))

	timestamp := time.Now().Unix()
	randomHash := genRandomHash(randomNumber)

	ethSwapID, err := e.calcSwapId(randomHash, ethSenderAddr, emSenderAddr)
	assertNoError(err, "e.calcSwapId()")

	trxHash, err := e.crEthTrx(timestamp, randomHash)
	assertNoError(err, "HTLT creation failed")
	fmt.Println("created trx hash:", trxHash)
	fmt.Println("sleeping 2 seconds...")

	time.Sleep(2 * time.Second)

	exists, err := e.hasSwap(common.BytesToHash(ethSwapID))
	assertNoError(err, "e.hasSwap()")
	fmt.Printf("query HasSwap() returns %t, for swapID:%s\n", exists, hex.EncodeToString(ethSwapID))

	swap, err := e.getSwap(common.BytesToHash(ethSwapID))
	assertNoError(err, "e.getSwap()")

	fmt.Println("getSwap() retuns -> swap obj:")
	fmt.Println("swap.Id", swap.Id.String())
	fmt.Println("swap.RandomNumberHash", swap.RandomNumberHash.String())
	fmt.Println("swap.ExpireTimestamp", swap.ExpireTimestamp)
	fmt.Println("swap.OutAmount", swap.OutAmount.String())
	fmt.Println("swap.RecipientAddress", swap.RecipientAddress)
	fmt.Println("swap.RecipientOtherChain", swap.RecipientOtherChain)
	fmt.Println("swap.SenderAddress", swap.SenderAddress)

	return nil
}

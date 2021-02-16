package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	emtypes "github.com/e-money/em-ledger/types"
	"github.com/ethereum/go-ethereum/common"
	et "github.com/ethereum/go-ethereum/core/types"
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

	emtypes.SetNetworkPrefix(sdk.GetConfig())

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

	time.Sleep(2 * time.Second)

	exists, err := e.hasSwap(common.BytesToHash(ethSwapID))
	assertNoError(err, "e.hasSwap()")
	fmt.Printf("query eth HasSwap exists %t, swapID:%s\n", exists, hex.EncodeToString(ethSwapID))

	swap, err := e.getSwap(common.BytesToHash(ethSwapID))
	assertNoError(err, "e.getSwap()")
	fmt.Printf("getSwap() error:%v\n", swap)

	trxStatus, blockNum, err := e.getTrxReceiptStatus(trxHash)
	assertNoError(err, "e.getTrxReceiptStatus()")

	fmt.Printf("height:%d, trx status:%d, suceeded:%t\n", blockNum, trxStatus, trxStatus == et.ReceiptStatusSuccessful)

	ethBal, err := e.ethBalance(common.HexToAddress(ethSenderAddr))
	assertNoError(err, "ethBalance(ethSender)")
	fmt.Printf("query eth balance:%d of %s\n", ethBal, ethSenderAddr)

	erc20Allowance, err := e.allowance()
	assertNoError(err, "allowance()")
	fmt.Printf("query allowance:%d\n", erc20Allowance)

	erc20Bal, err := e.erc20Balance(common.HexToAddress(ethSenderAddr))
	assertNoError(err, "erc20Balance()")
	fmt.Printf("query erc20Balance:%d of %s\n", erc20Bal, ethSenderAddr)

	blockAndTxLogs, err := e.getBlockAndTxs(blockNum.Int64())
	assertNoError(err, fmt.Sprintf("getblockAndTxs error%s, height=%d", err, blockNum))
	fmt.Println("blockLogs:", blockAndTxLogs)

	exists, err = e.hasSwap(common.BytesToHash(ethSwapID))
	assertNoError(err, "hasSwap() 2nd time")
	fmt.Printf("HasSwap():%t\n", exists)

	return nil
}
